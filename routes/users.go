package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	err = user.SaveUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Couldn't add user. Try again later.",
		})
		return
	}

	context.JSON(http.StatusCreated, models.UsersResponse{
		StatusCode: http.StatusCreated,
		Message:    http.StatusText(http.StatusCreated),
	})
}

// This function purpose for show all users that already registered
func getUsers(context *gin.Context) {
	users, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Couldn't fetch users. Try again later.",
		})
		return
	}

	response := models.UsersResponse{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		User:       users,
	}

	context.JSON(http.StatusOK, response)
}
