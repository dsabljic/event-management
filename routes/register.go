package routes

import (
	"strconv"

	"github.com/dsabljic/event-management/models"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(404, gin.H{"error": "Could not retrieve event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not register for event"})
		return
	}

	context.JSON(200, gin.H{"message": "Successfully registered for event"})
}

func unregister(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(404, gin.H{"error": "Could not retrieve event"})
		return
	}

	// var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not unregister for event"})
		return
	}

	context.JSON(200, gin.H{"message": "Successfully unregistered for event"})
}

func getEventRegistrations(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event ID"})
		return
	}

	registrations, err := models.FetchRegistrations(eventId)
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not fetch registrations"})
		return
	}

	context.JSON(200, gin.H{"registrations": registrations})
}
