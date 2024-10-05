package controllers

import (
	"database/sql"
	"fmt"
	"gfreecs0510/events/src/models"
	"gfreecs0510/events/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)
	if err != nil {
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

	event.UserId = userId

	err = event.Create()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func UpdateEvent(context *gin.Context) {
	var eventForUpdate models.Event
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse event id",
		})
		return
	}

	err = context.ShouldBind(&eventForUpdate)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot bind event body",
		})
		return
	}

	eventForUpdate.ID = id

	eventExists, err := models.GetEventViaId(id)

	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "event not found",
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	userId, err := utils.GetAuthenticatedUserId(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if eventExists.UserId != userId {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "unauthorized",
		})
	}

	err = eventForUpdate.Update()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "update event error",
		})
		return
	}

	context.JSON(http.StatusOK, eventForUpdate)
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse event id",
		})
		return
	}

	eventForDelete, err := models.GetEventViaId(id)

	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "event not found",
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	userId, err := utils.GetAuthenticatedUserId(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if eventForDelete.UserId != userId {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "unauthorized",
		})
	}

	err = eventForDelete.Delete()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete event error",
		})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "event deleted",
	})
}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEventViaId(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse event id",
		})
		return
	}

	event, err := models.GetEventViaId(id)

	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "event not found",
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	context.JSON(http.StatusOK, event)
}
