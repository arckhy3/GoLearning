package routes

import (
	"net/http"
	"strconv"

	"example.com/event/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	var register = models.Register{}

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse data", "error": err.Error()})
		return
	}

	validEvent, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to find event", "error": err.Error()})
		return
	}

	userID := context.GetInt64("userID")

	register.UserID = userID
	register.EventID = validEvent.ID

	err = register.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create registration", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration success", "data": register})
}

func cancelRegistration(context *gin.Context) {
	var register = models.Register{}

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse data", "error": err.Error()})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data", "error": err.Error()})
		return
	}

	userID := context.GetInt64("userID")
	register.UserID = userID
	register.EventID = eventID

	err = register.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unregister", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success unregister"})
}
