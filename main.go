package main

import (
	"net/http"

	"example.com/go-rest-api/db"
	"example.com/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event                // declare a variable of type Event from models package
	err := context.ShouldBindJSON(&event) // bind the JSON data from the request body to the event variable, and check for errors
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass request data.", "error": err.Error()})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save() // save the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
