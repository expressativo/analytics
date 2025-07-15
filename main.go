package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"analytics/config"
	"analytics/models"
	"analytics/routes"
)

func main() {
	// Inicializar la base de datos
	config.InitDB()

	// Migrar el esquema
	config.DB.AutoMigrate(&models.Event{})

	// Configurar el router
	r := gin.Default()

	// Configurar las rutas
	routes.SetupEventRoutes(r)

	// Obtener el puerto del servidor desde las variables de entorno
	serverPort := config.GetEnv("SERVER_PORT", "8080")

	// Iniciar el servidor
	if err := r.Run(":" + serverPort); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
