package routes

import (
	"github.com/dsabljic/event-management/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.GET("/events/:id/registrations", getEventRegistrations)
	authenticated.POST("/events/:id/register", register)
	authenticated.DELETE("/events/:id/register", unregister)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
