package config

import "os"

// GetEnv obtiene una variable de entorno o devuelve un valor por defecto
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
