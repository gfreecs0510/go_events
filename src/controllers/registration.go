package controllers

import (
	"fmt"
	"gfreecs0510/events/src/models"
	"gfreecs0510/events/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func CreateRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	userId, err := utils.GetAuthenticatedUserId(context)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	var reg models.Registration

	reg.UserId = userId
	reg.EventId = eventId

	ifExists, _ := models.GetEventViaId(eventId)

	if ifExists == (models.Event{}) {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "event does not exists",
		})
		return
	}

	err = reg.Create()

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				context.JSON(http.StatusBadRequest, gin.H{
					"message": "already registered",
				})
				return
			}
		}

		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, reg)
}

func DeleteRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	userId, err := utils.GetAuthenticatedUserId(context)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	var reg models.Registration

	reg.UserId = userId
	reg.EventId = eventId

	err = reg.Delete()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "registration deleted",
	})
}
