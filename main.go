package main

import (
	"filesharing/share/handlers/authentication"
	"filesharing/share/handlers/service"
	"filesharing/share/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.POST("/auth/signup", authentication.Signup)
	router.GET("/packPock/home", middlewares.AuthorizeMiddle, service.HomepageService)
	router.Run(":8000")
}
