package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Save() (any, error)
}

func RegisterRoute(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", saveEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
