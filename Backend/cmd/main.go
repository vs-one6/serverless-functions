package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vs-one6/serverless-functions.git/Backend/Routes"
)

func main() {
	//Load env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found ")
	}
	server := gin.Default()
	log.Println("Server starting.....")
	Routes.RegisterRoutes(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback port
	}
	log.Println("Server running on port ", port)

	server.Run(":" + port)
}
