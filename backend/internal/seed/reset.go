package seed
import (
	"log"
	"gorm.io/gorm"
)
func Reset(db *gorm.DB) error {
	if err := db.Exec("DELETE FROM orders").Error; err != nil {
		return err
	}
	log.Println("✅ Deleted all orders")
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	log.Println("✅ Deleted all users")
	if err := db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1").Error; err != nil {
		log.Printf("Warning: Failed to reset users sequence: %v", err)
	}
	if err := db.Exec("ALTER SEQUENCE orders_id_seq RESTART WITH 1").Error; err != nil {
		log.Printf("Warning: Failed to reset orders sequence: %v", err)
	}
	return nil
}
