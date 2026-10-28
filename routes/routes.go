package routes

import (
	"example.com/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
	server.GET("/events/:id", getEvent) // e.g. /events/1, /events/2, etc.

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // apply the authentication middleware to the group;
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent) // /events/1 etc.

	server.POST("/signup", createUser)
	server.POST("/login", login) // /login endpoint for user authentication
}
