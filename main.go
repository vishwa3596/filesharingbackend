package main

import (
	"filesharing/share/handlers"
	"filesharing/share/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.POST("/auth/login", handlers.Login)
    router.POST("/auth/signup", handlers.Signup)
	router.Run(":8000")
}
