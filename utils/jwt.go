package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email string,userId int64) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64,error) {
	parsedToken, err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		_,ok:=t.Method.(*jwt.SigningMethodHMAC)//(SigningMethodHS256) is a version of this (SigningMethodHMAC)
		
		if !ok {
			return nil, errors.New("unexpected signing method")	
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0,errors.New("invalid Token")
	}

	claims,ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
 return userId, nil
}