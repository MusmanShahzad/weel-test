package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	OpenAI   OpenAIConfig
	JWT      JWTConfig
}
type ServerConfig struct {
	Port string
	Host string
}
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}
type OpenAIConfig struct {
	Secret string
}
type JWTConfig struct {
	Secret string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "weel_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		OpenAI: OpenAIConfig{
			Secret: getEnv("OPEN_AI_SECRET", ""),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "default-secret-change-in-production"),
		},
	}
	return config, nil
}
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
