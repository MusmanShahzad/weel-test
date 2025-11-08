package user
import (
	"weel-backend/internal/handler"
	"weel-backend/internal/module"
	"weel-backend/internal/repository"
	"weel-backend/internal/router"
	"weel-backend/internal/service"
	"gorm.io/gorm"
)
type UserModule struct {
	userRepo    repository.UserRepository
	userService service.UserService
	userHandler *handler.UserHandler
}
func NewUserModule() module.Module {
	return &UserModule{}
}
func (m *UserModule) Name() string {
	return "user"
}
func (m *UserModule) Initialize(db *gorm.DB) error {
	m.userRepo = repository.NewUserRepository(db)
	m.userService = service.NewUserService(m.userRepo)
	m.userHandler = handler.NewUserHandler(m.userService)
	return nil
}
func (m *UserModule) RegisterRoutes(r *router.Router) {
	r.RegisterRoutes(m.userHandler)
}
