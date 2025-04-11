package routes

import (
	"example.com/event/middlewares"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Save() (any, error)
	UpdateByID() error
	Delete() error
}

func RegisterRoute(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//required token when called
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", saveEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)

	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("events/:id/register", cancelRegistration)

	//uncomment if you want
	// server.GET("/users", getUsers)
	// server.GET("/users/:id", getUser)
	// server.PUT("/user/:id", updateUser)
	// server.DELETE("/user/:id", deleteUser)

	server.POST("/signup", saveUser)
	server.POST("/login", login)
}
