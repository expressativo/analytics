package routes

import (
	"github.com/gin-gonic/gin"

	"analytics/controllers"
)

// SetupEventRoutes configura las rutas para los eventos
func SetupEventRoutes(router *gin.Engine) {
	eventController := controllers.NewEventController()

	// Rutas para el CRUD de eventos
	router.POST("/events", eventController.CreateEvent)
	router.GET("/events", eventController.GetEvents)
	router.GET("/events/:id", eventController.GetEvent)
	// router.PUT("/events/:id", eventController.UpdateEvent)
	// router.DELETE("/events/:id", eventController.DeleteEvent)
}
