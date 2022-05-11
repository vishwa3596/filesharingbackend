package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Gethello(c *gin.Context) {
	fmt.Println("HI")
}
