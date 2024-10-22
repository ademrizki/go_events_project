package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEventID)
	server.DELETE("/events/:id", deleteEvent)
}
