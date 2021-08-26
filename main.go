package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	router.GET("/", helloWorld)
	router.POST("/send-message", sendMessage)

	router.Run("localhost:8000")
}

func helloWorld(c *gin.Context) {
	greeting := map[string]string{
		"message": "Welcome to the https://coped.dev API!",
	}
	c.IndentedJSON(http.StatusOK, greeting)
}

func sendMessage(c *gin.Context) {
	var request message
	if err := c.BindJSON(&request); err != nil {
		return
	}
}
