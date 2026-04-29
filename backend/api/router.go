package api

import (
	"ai-wordbook/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", Register)
		api.POST("/login", Login)

		authorized := api.Group("")
		authorized.Use(middleware.JWTAuth())
		{
			authorized.POST("/words/query", QueryWord)
			authorized.POST("/words", SaveWord)
			authorized.GET("/words", ListWords)
			authorized.DELETE("/words/:id", DeleteWord)
		}
	}
}
