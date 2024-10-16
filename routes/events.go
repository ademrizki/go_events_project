package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Couldn't fetch events. Try again later.",
		})
		return
	}

	response := models.EventsResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Event:      events,
	}

	context.JSON(http.StatusOK, response)

}

func getEventByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Parse ID failed.",
		})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "There isn't any event for this ID",
		})
		return
	}

	response := models.EventsResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Event:      &event,
	}

	context.JSON(http.StatusOK, response)
}

func createEvents(context *gin.Context) {
	var events models.Event
	err := context.ShouldBindJSON(&events)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	events.UserID = 1

	err = events.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Couldn't create event. Try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, models.EventsResponse{
		StatusCode: http.StatusCreated,
		Message:    http.StatusText(http.StatusCreated),
		Event:      events,
	})
}
