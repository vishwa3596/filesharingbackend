package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserSignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signup(ctx *gin.Context) {
	var currentUser UserSignUp
	errorFromBinding := ctx.BindJSON(&currentUser)

	if errorFromBinding != nil {
		ctx.JSON(http.StatusBadRequest, "Error Binding Json")
		return
	}

	jwtSigningKey := []byte(currentUser.Password)

	claims := UserClaim{
		currentUser.Username,
		jwt.StandardClaims{
			ExpiresAt: 1500,
			Issuer:    "PackPock",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSigningKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, " Error Creating Jwt Token ")
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"token": signedToken,
	})
}
