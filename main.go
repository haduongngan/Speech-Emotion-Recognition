package main

import (
	"log"
	"net/http"
	"os"
	"spser/infrastructure"
	"spser/router"
)

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "19001"
	}

	log.Fatal(http.ListenAndServe(":"+appPort, router.Router()))

}
