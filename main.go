package main

import (
	"net/http"

	"example.com/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080
}
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"events": events})
}
func createEvent(context *gin.Context) {
	var event models.Event // declare a variable of type Event from models package
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save() // save the event
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
