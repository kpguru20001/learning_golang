package configs

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  int
	WriteTimeout int
}

// DatabaseConfig holds database-specific configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	AuthDB   string
	SSLMode  string
}

// LoggingConfig holds logging-specific configuration
type LoggingConfig struct {
	Level string
}

var Configuration *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	Configuration = &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			Host:         getEnv("HOST", "localhost"),
			ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 15),
			WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 15),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "27017"),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "shadow_docs"),
			AuthDB:   getEnv("DB_AUTH_DB", "admin"),
			SSLMode:  getEnv("DB_SSL_MODE", "false"),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}

	return Configuration
}

func getEnv(key string, defaultValue string) string {

	value := viper.GetString(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {

	value := viper.GetString(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting %s to int, using default value %d: %v", key, defaultValue, err)
		return defaultValue
	}

	return intValue
}
