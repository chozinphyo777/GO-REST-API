package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
	server.GET("/events/:id", getEvent) // e.g. /events/1, /events/2, etc.
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent) // /events/1 etc.
}
