package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "dummysupersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parseToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signin method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	validToken := parseToken.Valid

	if !validToken {
		return 0, errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
