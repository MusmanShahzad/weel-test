package database
import (
	"fmt"
	"log"
	"weel-backend/config"
	"weel-backend/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB
func Connect(cfg *config.Config) error {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Println("Database connected successfully")
	return nil
}
func AutoMigrate() error {
	err := DB.AutoMigrate(
		&domain.User{},
		&domain.Order{},
		&domain.FeatureFlag{},
	)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	log.Println("Database migrations completed successfully")
	return nil
}
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
