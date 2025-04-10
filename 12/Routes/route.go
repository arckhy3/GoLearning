package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Save() (any, error)
	UpdateByID() error
	Delete() error
}

func RegisterRoute(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", saveEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.PUT("/user/:id", updateUser)
	server.DELETE("/user/:id", deleteUser)

	server.POST("/signup", saveUser)
	server.POST("/login", login)
}
