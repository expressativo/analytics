package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"analytics/config"
	"analytics/models"
)

// EventController maneja las operaciones CRUD para eventos
type EventController struct {
	DB *gorm.DB
}

// NewEventController crea una nueva instancia del controlador de eventos
func NewEventController() *EventController {
	return &EventController{DB: config.DB}
}

// CreateEvent crea un nuevo evento
func (ec *EventController) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentTime := time.Now().Format(time.RFC3339)
	event.Date = currentTime

	if event.Metadata == "" {
		event.Metadata = "{}"
	}

	result := ec.DB.Create(&event)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

// GetEvents obtiene todos los eventos
func (ec *EventController) GetEvents(c *gin.Context) {
	var events []models.Event
	result := ec.DB.Find(&events)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetEvent obtiene un evento por su ID
func (ec *EventController) GetEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	result := ec.DB.First(&event, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// UpdateEvent actualiza un evento existente
func (ec *EventController) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	result := ec.DB.First(&event, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ec.DB.Save(&event)
	c.JSON(http.StatusOK, event)
}

// DeleteEvent elimina un evento
func (ec *EventController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	result := ec.DB.First(&event, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	ec.DB.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Evento eliminado correctamente"})
}
