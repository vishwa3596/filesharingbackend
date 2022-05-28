package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserRefreshClaim struct {
	UniqueId [16]byte `json:"uuid"`
	jwt.StandardClaims
}

func GenerateKey(userName string) (map[string]string, error) {
	secret := []byte("ASDSADWE##@LKSDJFLDKJ#R#$@#$#EWRSLFJSDLFKJSDFLKEWJRLWKEJRL23123234@##")

	claims := UserClaim{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "PackPock",
		},
	}

	refreshId, _ := uuid.NewRandom()

	refreshClaims := UserRefreshClaim{
		refreshId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "PackPock",
			IssuedAt:  time.Now().Unix(),
		},
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessToken, errorAccessToken := aToken.SignedString(secret)

	if errorAccessToken != nil {
		return map[string]string{
			"accessToken":  "",
			"refreshToken": "",
		}, errors.New("error signing access token")
	}
	refreshToken, errorRefreshToken := rToken.SignedString(secret)

	if errorRefreshToken != nil {
		return map[string]string{
			"accessToken":  "",
			"refreshToken": "",
		}, errors.New("error signing refresh token")
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil

}
