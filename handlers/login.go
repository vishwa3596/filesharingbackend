package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

type userClaim struct {
	firstName      string `json:"firstname"`
	lastName       string `json:"lastname"`
	password       string `json:"password"`
	uniqueUsername string `json:"username"`
	jwt.StandardClaims
}

func Login(ctx *gin.Context) {

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := userClaim{
		"bar",
		"bar",
		"bar",
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, " Error Signing String ")
		return
	}
	ctx.JSON(http.StatusAccepted, ss)
}
