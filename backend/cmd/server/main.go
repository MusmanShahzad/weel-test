package main
import (
	"fmt"
	"log"
	"net/http"
	"weel-backend/config"
	"weel-backend/internal/app"
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
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	defer application.Close()
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, application.GetRouter().Router.GetEngine()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
