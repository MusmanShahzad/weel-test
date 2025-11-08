package main
import (
	"log"
	"weel-backend/config"
	"weel-backend/internal/database"
)
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()
	if err := database.AutoMigrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("✅ Migrations completed successfully")
}
