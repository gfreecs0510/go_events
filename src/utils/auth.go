package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type AuthenticatedUser struct {
	ID       int64
	Username string
}

func GetAuthenticatedUser(context *gin.Context) (AuthenticatedUser, error) {
	var authUser AuthenticatedUser
	auIn, exists := context.Get("authenticated_user")

	if !exists {
		return authUser, errors.New("authenticated_user not in context")
	}

	authUser, ok := auIn.(AuthenticatedUser)

	if !ok {
		return authUser, errors.New("failed to assert authenticated_user")
	}

	return authUser, nil
}

func GetAuthenticatedUserId(context *gin.Context) (int64, error) {
	authUser, err := GetAuthenticatedUser(context)
	return authUser.ID, err
}
