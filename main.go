package main

import (
	"coped.dev/api/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	v1 := router.Group("/v1")
	{

		v1.GET("/", handlers.Index)
		v1.POST("/contact", handlers.Contact)

	}

	router.Run("localhost:8000")
}
