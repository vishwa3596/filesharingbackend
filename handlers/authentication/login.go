package authentication

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserLoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(ctx *gin.Context) {

	var userDetails UserLoginDetails

	err := ctx.BindJSON(&userDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding data to json during login",
		})
		return
	}

	loginSigningKey := []byte(userDetails.Password)

	// Create the Claims
	loginClaims := UserLoginClaim{
		userDetails.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 20).Unix(),
			Issuer:    "PackPock",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, loginClaims)
	accessToken, errorSigningTokens := token.SignedString(loginSigningKey)
	if errorSigningTokens != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Signing token during Login",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"access_token": accessToken,
	})
}
