package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.naous.net/api/models"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch particula event. Missing ID."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot get data for event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch events"})
		return
	}
	context.JSON(http.StatusOK, events) // context.JSON(http.StatusOK, gin.H{"message": "Hello", "command": "Obey"})

}

func createEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse request"})
		return
	}

	userId := context.GetInt64("userid")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch particula event. Missing ID."})
		return
	}
	userId := context.GetInt64("userid")
	event, err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot update! This event belongs to another user."})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot get data for event!"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse request!"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot update this event!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot fetch particula event. Missing ID."})
		return
	}
	userId := context.GetInt64("userid")
	event, err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot update! This event belongs to another user."})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot get data for event"})
		return
	}

	var deletedEvent models.Event

	// err = context.ShouldBindJSON(&deletedEvent)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse request"})
	// 	return
	// }

	deletedEvent.ID = eventId
	err = deletedEvent.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot delete this event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted", "event": deletedEvent})
}
