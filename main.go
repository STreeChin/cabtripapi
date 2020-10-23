package main

import (
	"github.com/STreeChin/cabtripapidomain/model"
	routers "github.com/STreeChin/cabtripapi/route"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	router := routers.NewRouter()

	model.MongoDBConnect()
	defer model.MongoConnectionClose()

	//MySQL
	//model.MySQLOpen()
	//model swagger.MySQLClose()

	log.Fatal(http.ListenAndServe(":8080", router))
}
