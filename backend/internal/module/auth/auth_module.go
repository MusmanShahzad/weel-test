package auth

import (
	"weel-backend/internal/handler"
	"weel-backend/internal/middleware"
	"weel-backend/internal/module"
	"weel-backend/internal/repository"
	"weel-backend/internal/router"
	"weel-backend/internal/service"

	"gorm.io/gorm"
)

type AuthModule struct {
	userRepo    repository.UserRepository
	jwtService  *service.JWTService
	authService service.AuthService
	authHandler *handler.AuthHandler
}

func NewAuthModule() module.Module {
	return &AuthModule{}
}
func (m *AuthModule) Name() string {
	return "auth"
}
func (m *AuthModule) Initialize(db *gorm.DB) error {
	m.userRepo = repository.NewUserRepository(db)
	m.jwtService = service.NewJWTService()
	m.authService = service.NewAuthService(m.userRepo, m.jwtService)
	m.authHandler = handler.NewAuthHandler(m.authService)
	return nil
}
func (m *AuthModule) RegisterRoutes(r *router.Router) {
	v1 := r.GetEngine().Group("/api/v1")
	v1.POST("/auth/login", m.authHandler.Login)
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware(m.jwtService))
	protected.GET("/me", m.authHandler.GetMe)
}
