package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/event/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get data.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func saveEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	userID := context.GetInt64("userID")
	event.UserID = userID
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Data failed to create.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get data.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data", "error": err.Error()})
		return
	}

	var updatedEvent = models.Event{}
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	updatedEvent.ID = id
	userID := context.GetInt64("userID")

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to update event"})
		return
	}

	updatedEvent.UserID = userID
	err = updatedEvent.UpdateByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update data", "error": err.Error()})
		return
	}
	fmt.Println(updatedEvent)
	context.JSON(http.StatusOK, gin.H{"message": "Update success!", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get data.", "error": err.Error()})
		return
	}
	e, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data", "error": err.Error()})
		return
	}
	userID := context.GetInt64("userID")

	if e.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorize to delete event"})
		return
	}

	err = e.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete data", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete success"})
}
