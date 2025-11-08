package service
import (
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
)
type FeatureFlagService interface {
	GetAllFlags() ([]*domain.FeatureFlag, error)
	GetFlagByName(name string) (*domain.FeatureFlag, error)
	IsEnabled(name string) bool
	UpdateFlag(name string, enabled bool) (*domain.FeatureFlag, error)
}
type featureFlagService struct {
	flagRepo repository.FeatureFlagRepository
}
func NewFeatureFlagService(flagRepo repository.FeatureFlagRepository) FeatureFlagService {
	return &featureFlagService{flagRepo: flagRepo}
}
func (s *featureFlagService) GetAllFlags() ([]*domain.FeatureFlag, error) {
	return s.flagRepo.GetAll()
}
func (s *featureFlagService) GetFlagByName(name string) (*domain.FeatureFlag, error) {
	return s.flagRepo.GetByName(name)
}
func (s *featureFlagService) IsEnabled(name string) bool {
	flag, err := s.flagRepo.GetByName(name)
	if err != nil {
		return false
	}
	return flag.Enabled
}
func (s *featureFlagService) UpdateFlag(name string, enabled bool) (*domain.FeatureFlag, error) {
	flag, err := s.flagRepo.GetByName(name)
	if err != nil {
		return nil, ErrUserNotFound
	}
	flag.Enabled = enabled
	if err := s.flagRepo.Update(flag); err != nil {
		return nil, err
	}
	return flag, nil
}
