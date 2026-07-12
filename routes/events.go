package routes

import (
	"net/http"
	"strconv"

	"example.com/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id parameter from the URL
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})

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
