package cabsvc

import (
	"github.com/STreeChin/cabtripapi/domain/model"
	"context"
	"net/http"
	"strconv"
	"strings"
)

func GetCabTrip(w http.ResponseWriter, cab, date []string, fresh bool) {
	//the parameters have been checked in Ctrl function, dont need defence programming here.
	dbCount, cacheCount, err := int64(0), "", error(nil)
	var rsp strings.Builder

	db := model.GetMongodbClient()
	cache := model.GetCacheInstance()
	//fresh does not exist, cache first
	if fresh == false {
		for i, _ := range cab {
			qdate := date[0]
			if len(date) != 1 {
				qdate = date[i]
			}
			//get cache first, if missing read DB
			if cacheCount, err = cache.GetCache(cab[i] + "_" + qdate); err != nil {
				//For those cab in date whose trips is 0, dont set cache
				if dbCount = db.GetCount(context.Background(), cab[i], qdate); dbCount != 0 {
					cache.SetCache(cab[i]+"_"+qdate, strconv.FormatInt(dbCount, 10))
				}
				rsp.WriteString(cab[i] + " in " + qdate + ":" + strconv.FormatInt(dbCount, 10) + "\n")
			} else {
				rsp.WriteString(cab[i] + " in " + qdate + ":" + cacheCount + "\n")
			}
		}
	} else {
		for i, _ := range cab {
			qdate := date[0]
			if len(date) != 1 {
				qdate = date[i]
			}
			//For those cab in date whose trips is 0, dont set cache
			if dbCount = db.GetCount(context.Background(), cab[i], qdate); dbCount != 0 {
				//if count in cache equals count from db, dont need to set cache
				if cacheCount, _ = cache.GetCache(cab[i] + "_" + qdate); cacheCount != strconv.FormatInt(dbCount,10) {
					cache.SetCache(cab[i]+"_"+qdate, strconv.FormatInt(dbCount,10))
				}
			}
			rsp.WriteString(cab[i] + " in " + qdate + ":" + strconv.FormatInt(dbCount,10) + "\n")
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rsp.String()))
}
