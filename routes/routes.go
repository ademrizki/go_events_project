package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", ValidationMiddleware(), getEvents)
	server.GET("/events/:id", ValidationMiddleware(), getEventByID)
	server.POST("/events", ValidationMiddleware(), createEvents)
	server.PUT("/events/:id", ValidationMiddleware(), updateEventID)
	server.DELETE("/events/:id", ValidationMiddleware(), deleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/users", ValidationMiddleware(), getUsers)
}

func ValidationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		token := context.GetHeader("Authorization")

		if token == "" {
			context.JSON(http.StatusUnauthorized, models.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token doesn't exist",
			})
			context.Abort()
			return
		}

		tokenString := token[len(BEARER_SCHEMA):]

		userID, err := utils.ValidateToken(tokenString)

		if err != nil {
			context.JSON(http.StatusUnauthorized, models.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
			})
			context.Abort()
			return
		}

		if userID == 0 {
			context.JSON(http.StatusUnauthorized, models.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "User doesn't exist",
			})
			context.Abort()
			return
		}

		context.Set("user_id", userID)
		context.Next()
	}

}
