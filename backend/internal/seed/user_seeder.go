package seed
import (
	"log"
	"weel-backend/internal/domain"
	"weel-backend/internal/service"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)
type UserSeeder struct{}
func NewUserSeeder() Seeder {
	return &UserSeeder{}
}
func (s *UserSeeder) Name() string {
	return "UserSeeder"
}
func (s *UserSeeder) Seed(db *gorm.DB) error {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		log.Println("Users already exist, skipping seed")
		return nil
	}
	adminPassword, _ := service.HashPassword("password123")
	admin := &domain.User{
		Email:     "admin@example.com",
		Password:  adminPassword,
		FirstName: "Admin",
		LastName:  "User",
	}
	if err := db.Create(admin).Error; err != nil {
		return err
	}
	log.Printf("✅ Created admin user: %s", admin.Email)
	testPassword, _ := service.HashPassword("password123")
	testUser := &domain.User{
		Email:     "user@example.com",
		Password:  testPassword,
		FirstName: "Test",
		LastName:  "User",
	}
	if err := db.Create(testUser).Error; err != nil {
		return err
	}
	log.Printf("✅ Created test user: %s", testUser.Email)
	fakeUsers := make([]*domain.User, 10)
	for i := 0; i < 10; i++ {
		password, _ := service.HashPassword("password123")
		fakeUsers[i] = &domain.User{
			Email:     gofakeit.Email(),
			Password:  password,
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
		}
	}
	if err := db.CreateInBatches(fakeUsers, 10).Error; err != nil {
		return err
	}
	log.Printf("✅ Created %d fake users", len(fakeUsers))
	return nil
}
