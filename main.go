package main

import (
	"log"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router.Router()))

}
