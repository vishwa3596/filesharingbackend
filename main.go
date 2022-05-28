package main

import (
	"filesharing/share/handlers/authentication"
	"filesharing/share/handlers/service"
	"filesharing/share/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.POST("/auth/signup", authentication.Signup)
	router.POST("/auth/login", authentication.Login)
	router.GET("/packPock/home", middlewares.AuthorizeMiddle, service.HomepageService)
	log.Fatalln(router.Run(":8000"))
}
