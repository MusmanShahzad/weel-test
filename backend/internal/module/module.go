package module
import (
	"gorm.io/gorm"
	"weel-backend/internal/router"
)
type Module interface {
	Initialize(db *gorm.DB) error
	RegisterRoutes(router *router.Router)
	Name() string
}
