package main

import (
	"log"

	"ai-wordbook/api"
	"ai-wordbook/config"
	"ai-wordbook/model"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	model.InitDB()

	r := gin.Default()
	api.SetupRoutes(r)

	port := config.AppConfig.ServerPort
	log.Printf("Server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
