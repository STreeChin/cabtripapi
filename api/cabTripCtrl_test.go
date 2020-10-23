package cabapi_test

import (
	cabapi "github.com/STreeChin/cabtripapi/api"
	"github.com/STreeChin/cabtripapi/domain/model"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Use GOCONVEY test framework
func TestGetCabTripCtrl(t *testing.T) {
	model.MongoDBConnect()
	defer model.MongoConnectionClose()
	Convey("GetCabTripCtrl", func() {
		req, err := http.NewRequest("GET", "/api/cab/id3004672/date/2016-06-30?fresh=1", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		Convey("Normal: 1 cab", func() {

			cabapi.GetCabTripCtrl(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
			result, _ := ioutil.ReadAll(w.Result().Body)
			expct := "id3004672 in 2016-06-30:3"
			So(result, ShouldEqual, expct)
		})


	})
}