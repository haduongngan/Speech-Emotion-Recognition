package main

import (
	"log"
	"net/http"
	"spser/infrastructure"
	"spser/router"
)

// @title Speech Emotion Recognition APIs
// @version 1.0
// @description SPSER - Speech Emotion Recognition Service.

// @host localhost:19001
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))

}
