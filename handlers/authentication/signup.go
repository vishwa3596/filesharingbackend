package authentication

import (
	"filesharing/share/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserSignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func Signup(ctx *gin.Context) {
	var currentUser UserSignUp
	errorFromBinding := ctx.BindJSON(&currentUser)

	if errorFromBinding != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding the json",
		})
		return
	}

	tokens, err := utils.GenerateKey(currentUser.Username)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"signupError": err,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"tokens": tokens,
	})

}
