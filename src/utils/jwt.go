package utils

import (
	"errors"
	"gfreecs0510/events/src/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateUserJWT(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	tokenP := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	return tokenP.SignedString([]byte(secret))
}

func TryVerifyThenParseToken(token string) (AuthenticatedUser, error) {
	var user AuthenticatedUser
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return user, err
	}

	if !parsedToken.Valid {
		return user, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return user, errors.New("invalid token claim")
	}

	username, ok := claims["username"].(string)

	if !ok {
		return user, errors.New("invalid username")
	}

	floatUserId, ok := claims["id"].(float64)

	if !ok {
		return user, errors.New("invalid user id")
	}

	userId := int64(floatUserId)

	user.Username = username
	user.ID = userId

	return user, nil
}
