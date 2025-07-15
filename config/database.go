package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: Archivo .env no encontrado")
	}

	// Configurar la cadena de conexión a la base de datos usando variables de entorno
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("DB_USER", "postgres")
	dbPassword := GetEnv("DB_PASSWORD", "postgres")
	dbName := GetEnv("DB_NAME", "events")
	dbSslMode := GetEnv("DB_SSL_MODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)

	// Conectar a la base de datos
	var err2 error
	DB, err2 = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err2)
	}

	log.Println("Conexión a la base de datos establecida correctamente")
}
