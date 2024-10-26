package routes

import (
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

// GET, POST, PUT, PATCH, DELETE, OPTIONS

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // /events/1 or /events/43

	authenticated := server.Group("/")
	authenticated.Use(utils.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
