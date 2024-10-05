package controllers

import (
	"fmt"
	"gfreecs0510/events/src/models"
	"gfreecs0510/events/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "username/password required",
		})
		return
	}

	ifExist, err := models.GetUserViaUsername(user.UserName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if ifExist != (models.User{}) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "username already exists",
		})
		return
	}

	hash, err := utils.GenerateHash(user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	user.Password = hash

	err = user.Create()

	if err != nil {
		fmt.Println(err)

		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				msg := fmt.Sprintf("user %v already exists", user.UserName)
				context.JSON(http.StatusBadRequest, gin.H{
					"message": msg,
				})
				return
			}
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully created",
	})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "username/password required",
		})
		return
	}

	ifExist, err := models.GetUserViaUsername(user.UserName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if ifExist == (models.User{}) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "username does not exists",
		})
		return
	}

	if !utils.CompareHashAndUserPassword(ifExist.Password, user.Password) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "password does not match",
		})
		return
	}

	token, err := utils.GenerateUserJWT(ifExist)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
