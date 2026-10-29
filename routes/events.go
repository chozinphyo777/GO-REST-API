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

	userId := context.GetInt64("userId") // get the userId from the context, which was set by the authentication middleware
	event.UserID = userId                // set the UserID of the event to the userId from the token
	err = event.Save()                   // save the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id parameter from the URL
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // bind the JSON data from the request body to the event variable, and check for errors
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass request data.", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	updatedEvent.UserID = userId // keep the authenticated owner ID; the request body should not overwrite it
	err = updatedEvent.Update()  // Update the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create update event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the id parameter from the URL
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID.", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Delete event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
