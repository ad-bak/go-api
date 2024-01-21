package routes

import (
	"example.com/books/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// group with middleware
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEventById)

	//users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
