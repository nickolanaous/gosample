package routes

import (
	"github.com/gin-gonic/gin"
	"go.naous.net/api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/rsvp", rsvpToEvent)
	authenticated.DELETE("/events/:id/rsvp", cancelRsvpToEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
