package cabsvc_test

import (
	"github.com/STreeChin/cabtripapi/domain/model"
	cabsvc "github.com/STreeChin/cabtripapi/domain/service"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
//Use GOCONVEY test framework
func TestGetCabTrip(t *testing.T) {
	model.MongoDBConnect()
	defer model.MongoConnectionClose()
	Convey("GetCabTrip", func() {
		w := httptest.NewRecorder()

		Convey("Normal: 1 cab", func() {
			///api/cab/id3004672/date/2016-06-30?fresh=0
			cab, date, fresh := []string{"id3004672"}, []string{"2016-06-30"}, true
			cabsvc.GetCabTrip(w,cab,date,fresh)

			//test the code
			So(w.Code, ShouldEqual, http.StatusOK)
			//test the body
			result, _ := ioutil.ReadAll(w.Result().Body)
			expct := "id3004672 in 2016-06-30:3"
			So(result, ShouldEqual, expct)
		})

	})
}
