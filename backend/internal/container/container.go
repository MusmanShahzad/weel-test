package container
import (
	"weel-backend/internal/module"
	"weel-backend/internal/router"
	"gorm.io/gorm"
)
type Container struct {
	DB      *gorm.DB
	Router  *router.Router
	Modules []module.Module
}
func NewContainer() *Container {
	return &Container{
		Modules: make([]module.Module, 0),
	}
}
func (c *Container) RegisterModule(m module.Module) {
	c.Modules = append(c.Modules, m)
}
func (c *Container) Initialize() error {
	c.Router = router.NewRouter()
	for _, m := range c.Modules {
		if err := m.Initialize(c.DB); err != nil {
			return err
		}
	}
	for _, m := range c.Modules {
		m.RegisterRoutes(c.Router)
	}
	return nil
}
