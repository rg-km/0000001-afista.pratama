package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func SignJwt(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(token)

	tokenStr, err := token.SignedString([]byte(secret))
	fmt.Println(tokenStr)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyJwt(token string, secret string) (map[string]interface{}, error) {
	jwToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwToken.Valid {
		return nil, fmt.Errorf("invalid authorization token")
	}
	return jwToken.Claims.(jwt.MapClaims), nil
}
