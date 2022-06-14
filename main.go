package main

import (
	"log"
	"net/http"
	"spser/infrastructure"
	"spser/router"
)

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))

}
