package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type Authorize struct {
	BearerToken string `header:"Authorization"`
}

func AuthorizeMiddle(ctx *gin.Context) {
	bearer := Authorize{}

	if errorBindingBearer := ctx.ShouldBindHeader(&bearer); errorBindingBearer != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Not able to bind header for authorization",
		})
		return
	}
	tokenHeader := strings.Split(bearer.BearerToken, "Bearer ")

	if len(tokenHeader) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty token field",
		})
		return
	}

	token, err := jwt.Parse(tokenHeader[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("ASDSADWE##@LKSDJFLDKJ#R#$@#$#EWRSLFJSDLFKJSDFLKEWJRLWKEJRL23123234@##"), nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.JSON(http.StatusAccepted, gin.H{
				"claims": claims,
			})
			return
		}
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			// Redirect to login page
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": " That is not even a token ",
			})
			ctx.Abort()
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// token is expired go to login or generate the new pair of token
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "token is expired",
			})
			ctx.Abort()
			return
		} else {
			// redirect to login.
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": " can't validated token ",
			})
			ctx.Abort()
			return
		}
	} else {
		// redirect to login
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": " Couldn't handle this token ",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
