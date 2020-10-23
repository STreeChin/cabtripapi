package cabapi

import (
	cabsvc "github.com/STreeChin/cabtripapi/domain/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetCabTripCtrl(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		log.Fatal("input nil")
	}
	vars := mux.Vars(r)
	cab, date := strings.Split(vars["id"], ","), strings.Split(vars["date"], ",")
	if check, info := checkParam(cab, date); !check {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(info))
		return
	}
	fresh := true
	if f := r.URL.Query().Get("fresh"); f == ""{
		fresh = false
	}

	cabsvc.GetCabTrip(w,cab,date,fresh)
}

func checkParam(cab, date []string) (bool, string) {
	check, info := true, ""
	if len(cab) == 0 || len(date) == 0 {
		check, info = false, "The number of dates or cabs must not equal 0."
	} else if len(date) != 1 && len(date) != len(date) {
		check, info = false, "The number of dates must equal 1 or equal the number of cabs."
	} else if len(date) > 10 {
		check, info = false, "The number of dates must be less than 11."
	} else{

	}

	for _, d := range date {
		if _, err := time.Parse("2006-01-02", d); err != nil {
			check, info = false, "Date format must be 2006-01-02."
			break
		}
	}
	return check, info
}
