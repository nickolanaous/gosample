package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.naous.net/api/models"
)

func rsvpToEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch this event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch event. Missing ID."})
	}

	err = event.Rsvp(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot register for this event. Something wrong happened."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event!"})

}

func cancelRsvpToEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch this event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch event. Missing ID."})
	}

	err = event.CancelRsvp(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot de-register for this event. Something wrong happened."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully de-registered for event!"})

}
