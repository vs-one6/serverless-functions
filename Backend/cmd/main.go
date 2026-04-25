package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vs-one6/serverless-functions.git/Backend/Routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found ")
	}
	server := gin.Default()
	log.Println("Server starting.....")
	Routes.RegisterRoutes(server)

	server.Run()
}
