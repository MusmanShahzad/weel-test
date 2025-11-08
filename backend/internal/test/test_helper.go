package test
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func GetTestDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
func CleanupDB(db *gorm.DB, tables ...string) error {
	for _, table := range tables {
		if err := db.Exec("TRUNCATE TABLE " + table + " RESTART IDENTITY CASCADE").Error; err != nil {
			return err
		}
	}
	return nil
}
