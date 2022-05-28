package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomepageService(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "homepage service is activated",
	})
	return
}
