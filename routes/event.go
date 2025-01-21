package routes

import (
	"strconv"

	"github.com/dsabljic/event-management/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"error": "Error getting events."})
		return
	}
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err := event.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating event."})
		return
	}

	context.JSON(201, gin.H{"message": "event created successfully"})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	context.JSON(200, gin.H{"event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid event ID"})
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		context.JSON(403, gin.H{"error": "You are not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse request data"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error updating event"})
		return
	}
	context.JSON(200, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid event ID"})
	}

	userID := context.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userID {
		context.JSON(403, gin.H{"error": "You are not authorized to delete this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error occured while trying to delete event"})
		return
	}

	context.JSON(200, gin.H{"message": "Event deleted successfully"})
}
