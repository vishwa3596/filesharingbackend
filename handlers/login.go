package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// field to keep in the signup form
/**
	usr = {
        "primary email":
		"password":
		"confirm password":
		"uniqueUsername": auto suggested or put input.
        "mySigningKey" : unique username along with the title
	}
**/

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func Login(ctx *gin.Context) {

	mySigningKey := []byte("AllYourBase")

	claims := jwt.MapClaims{
		"firstname": "john",
		"lastname":  "sam",
		"ExpiresAt": time.Now().Add(time.Hour * 2).Unix(),
		"IssuedAt":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, " Error Signing String ")
		return
	}
	ctx.JSON(http.StatusAccepted, ss)
}

func Signup(ctx *gin.Context) {
	// Sign up and this generate an access token based on the credential and then the token is stored in local storage
	// and as each request is send with passing the authorization token. When the authorization token is failed login is
	// again prompted for the user to login.
	var user User
	ctx.BindJSON(&user)
	finalUser := &User{
		Email:    "aaa@gmail.com",
		Password: "falksdfj",
		Username: "JD",
	}

	finalUserJson, err := json.Marshal(finalUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, " Json Marshall Failed ")
		return
	}

	ctx.JSON(http.StatusAccepted, string(finalUserJson))
}
