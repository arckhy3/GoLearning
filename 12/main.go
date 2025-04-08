package main

import (
	routes "example.com/event/Routes"
	"example.com/event/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoute(server)

	server.Run(":8080")
}
