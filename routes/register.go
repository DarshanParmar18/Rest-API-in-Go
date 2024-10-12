package routes

import (
	"net/http"
	"project/rest_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)


func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	// if event.UserID != userId {
	// 	context.JSON(http.StatusUnauthorized,gin.H{"message":"not authourized to update"})
	// 	return
	// }

	err = event.Register(userId)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}

func cancelRegistration(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registrations"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration Cancelled"})
}