package seed
import (
	"log"
	"weel-backend/internal/domain"
	"gorm.io/gorm"
)
type FeatureFlagSeeder struct{}
func NewFeatureFlagSeeder() Seeder {
	return &FeatureFlagSeeder{}
}
func (s *FeatureFlagSeeder) Name() string {
	return "FeatureFlagSeeder"
}
func (s *FeatureFlagSeeder) Seed(db *gorm.DB) error {
	var count int64
	db.Model(&domain.FeatureFlag{}).Count(&count)
	if count > 0 {
		log.Println("Feature flags already exist, skipping seed")
		return nil
	}
	flags := []*domain.FeatureFlag{
		{
			Name:        "ai_suggestions",
			Description: "Enable AI-powered product suggestions for pharmacy orders",
			Enabled:     true,
		},
		{
			Name:        "email_notifications",
			Description: "Enable email notifications for order updates",
			Enabled:     false,
		},
		{
			Name:        "sms_notifications",
			Description: "Enable SMS notifications for order updates",
			Enabled:     false,
		},
		{
			Name:        "advanced_filtering",
			Description: "Enable advanced filtering options in dashboard",
			Enabled:     true,
		},
	}
	for _, flag := range flags {
		if err := db.Create(flag).Error; err != nil {
			return err
		}
		log.Printf("✅ Created feature flag: %s (enabled: %v)", flag.Name, flag.Enabled)
	}
	return nil
}
