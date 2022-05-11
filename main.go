package main

import (
	"filesharing/share/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Gethello)
	router.POST("/login", handlers.Login)
	router.Run(":8000")
}
