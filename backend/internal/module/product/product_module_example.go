package product
/*
package product
import (
	"weel-backend/internal/handler"
	"weel-backend/internal/module"
	"weel-backend/internal/repository"
	"weel-backend/internal/router"
	"weel-backend/internal/service"
	"gorm.io/gorm"
)
type ProductModule struct {
	productRepo    repository.ProductRepository
	productService service.ProductService
	productHandler *handler.ProductHandler
}
func NewProductModule() module.Module {
	return &ProductModule{}
}
func (m *ProductModule) Name() string {
	return "product"
}
func (m *ProductModule) Initialize(db *gorm.DB) error {
	m.productRepo = repository.NewProductRepository(db)
	m.productService = service.NewProductService(m.productRepo)
	m.productHandler = handler.NewProductHandler(m.productService)
	return nil
}
func (m *ProductModule) RegisterRoutes(r *router.Router) {
	r.RegisterRoutes(m.productHandler)
}
*/
