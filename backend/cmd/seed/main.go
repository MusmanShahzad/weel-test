package main
import (
	"flag"
	"log"
	"weel-backend/config"
	"weel-backend/internal/database"
	"weel-backend/internal/seed"
)
func main() {
	reset := flag.Bool("reset", false, "Reset seed data (delete existing data before seeding)")
	flag.Parse()
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
	if *reset {
		log.Println("Resetting seed data...")
		if err := seed.Reset(database.DB); err != nil {
			log.Fatal("Failed to reset seed data:", err)
		}
	}
	seeders := []seed.Seeder{
		seed.NewFeatureFlagSeeder(),
		seed.NewUserSeeder(),
		seed.NewOrderSeeder(),
	}
	log.Println("Starting database seeding...")
	if err := seed.Run(database.DB, seeders...); err != nil {
		log.Fatal("Failed to seed database:", err)
	}
	log.Println("✅ Seed data created successfully!")
}
