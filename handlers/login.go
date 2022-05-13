package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// field to keep in the login form
/**
	usr = {
		"firstName": " ",
		"lastName": " ",
		"password": 8 digits smallcase uppercase special char numbers
		"confirm password":
		"uniqueUsername": auto suggested or put input.
		"phone number to be linked":
		"Or email to be linked":
		"or Google Authentication to be linked":
	}
**/

func Login(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		http.Error(ctx.Writer, " Error Reading Request Data ", http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusAccepted, data)
}
