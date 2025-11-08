package repository
import (
	"weel-backend/internal/domain"
	"gorm.io/gorm"
)
type FeatureFlagRepository interface {
	GetAll() ([]*domain.FeatureFlag, error)
	GetByName(name string) (*domain.FeatureFlag, error)
	Update(flag *domain.FeatureFlag) error
}
type featureFlagRepository struct {
	db *gorm.DB
}
func NewFeatureFlagRepository(db *gorm.DB) FeatureFlagRepository {
	return &featureFlagRepository{db: db}
}
func (r *featureFlagRepository) GetAll() ([]*domain.FeatureFlag, error) {
	var flags []*domain.FeatureFlag
	err := r.db.Order("name ASC").Find(&flags).Error
	return flags, err
}
func (r *featureFlagRepository) GetByName(name string) (*domain.FeatureFlag, error) {
	var flag domain.FeatureFlag
	err := r.db.Where("name = ?", name).First(&flag).Error
	if err != nil {
		return nil, err
	}
	return &flag, nil
}
func (r *featureFlagRepository) Update(flag *domain.FeatureFlag) error {
	return r.db.Save(flag).Error
}
