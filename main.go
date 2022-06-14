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

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router.Router()))

}
