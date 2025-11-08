package feature_flag
import (
	"weel-backend/internal/handler"
	"weel-backend/internal/module"
	"weel-backend/internal/repository"
	"weel-backend/internal/router"
	"weel-backend/internal/service"
	"gorm.io/gorm"
)
type FeatureFlagModule struct {
	flagRepo    repository.FeatureFlagRepository
	flagService service.FeatureFlagService
	flagHandler *handler.FeatureFlagHandler
}
func NewFeatureFlagModule() module.Module {
	return &FeatureFlagModule{}
}
func (m *FeatureFlagModule) Name() string {
	return "feature_flag"
}
func (m *FeatureFlagModule) Initialize(db *gorm.DB) error {
	m.flagRepo = repository.NewFeatureFlagRepository(db)
	m.flagService = service.NewFeatureFlagService(m.flagRepo)
	m.flagHandler = handler.NewFeatureFlagHandler(m.flagService)
	return nil
}
func (m *FeatureFlagModule) RegisterRoutes(r *router.Router) {
	v1 := r.GetEngine().Group("/api/v1")
	m.flagHandler.RegisterRoutes(v1)
}
