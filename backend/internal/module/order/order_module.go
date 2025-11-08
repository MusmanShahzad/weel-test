package order

import (
	"weel-backend/config"
	"weel-backend/internal/handler"
	"weel-backend/internal/middleware"
	"weel-backend/internal/module"
	"weel-backend/internal/repository"
	"weel-backend/internal/router"
	"weel-backend/internal/service"

	"gorm.io/gorm"
)

type OrderModule struct {
	orderRepo    repository.OrderRepository
	orderService service.OrderService
	orderHandler *handler.OrderHandler
	jwtService   *service.JWTService
	aiService    service.AIService
	cfg          *config.Config
}

func NewOrderModule(cfg *config.Config) module.Module {
	return &OrderModule{
		cfg: cfg,
	}
}
func (m *OrderModule) Name() string {
	return "order"
}
func (m *OrderModule) Initialize(db *gorm.DB) error {
	m.orderRepo = repository.NewOrderRepository(db)
	m.aiService = service.NewAIService(m.cfg)
	m.orderService = service.NewOrderService(m.orderRepo, m.aiService)
	m.orderHandler = handler.NewOrderHandler(m.orderService)
	m.jwtService = service.NewJWTService()
	return nil
}
func (m *OrderModule) RegisterRoutes(r *router.Router) {
	v1 := r.GetEngine().Group("/api/v1")
	v1.POST("/orders/suggestions", m.orderHandler.GetAISuggestions)
	r.RegisterProtectedRoutes(middleware.AuthMiddleware(m.jwtService), m.orderHandler)
}
