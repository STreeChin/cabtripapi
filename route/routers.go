package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
	cabapi "github.com/STreeChin/cabtripapi/api"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Queries     []string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Queries(route.Queries...).
			Name(route.Name).
			Handler(handler)
	}
	//another method to register route: use subrouter
	/*s := router.PathPrefix("/api").Subrouter()
	// "/api/"
	s.HandleFunc("/", Index)
	// "/api/cab/{id}/date/{date}"
	s.HandleFunc("/cab/{id}/date/{date}", cabtrip.GetCabTrip)
	// "/products/{key}/details"
	s.HandleFunc("/cab/{id}/date/{date}/fresh", cabtrip.GetCabTrip)
	*/
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Please input api parameter.API Map:\n",
		"/api/cab/id132/date/20160101, get 132 cab trips of 20160101",
		"/api/cab/id132/date/20160101?fresh=1, get fresh data of 132 cab trips of 20160101", )
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		[]string{},
		Index,
	},

	Route{
		"GetCabTripCtrl",
		strings.ToUpper("Get"),
		//127.0.0.1:8080/api/cab/id3004672/date/2016-06-30
		"/api/cab/{id}/date/{date}",
		[]string{},
		cabapi.GetCabTripCtrl,
	},
	Route{
		"GetCabTripCtrl",
		strings.ToUpper("Get"),
		//127.0.0.1:8080/api/cab/id3004672/date/2016-06-30?fresh=1
		"/api/cab/{id}/date/{date}",
		[]string{"fresh", "{fresh}"},
		cabapi.GetCabTripCtrl,
	},

	Route{
		"DeleteCache",
		strings.ToUpper("Delete"),
		"/api/clearcache",
		[]string{},
		cabapi.DeleteCacheCtrl,
	},
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
