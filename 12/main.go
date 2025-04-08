package main

import (
	"net/http"

	"example.com/event/db"
	"example.com/event/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", saveEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get data.", "error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func saveEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = 1
	event, err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Data failed to create.", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
