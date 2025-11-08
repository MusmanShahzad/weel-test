package seed
import (
	"gorm.io/gorm"
)
type Seeder interface {
	Seed(db *gorm.DB) error
	Name() string
}
func Run(db *gorm.DB, seeders ...Seeder) error {
	for _, seeder := range seeders {
		if err := seeder.Seed(db); err != nil {
			return err
		}
	}
	return nil
}
