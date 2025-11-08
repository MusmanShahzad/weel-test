package router
import (
	"net/http"
	"weel-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)
type RouteHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
}
type Router struct {
	engine         *gin.Engine
	authMiddleware gin.HandlerFunc
}
func NewRouter() *Router {
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	return &Router{
		engine: engine,
	}
}
func (r *Router) SetAuthMiddleware(middleware gin.HandlerFunc) {
	r.authMiddleware = middleware
}
func (r *Router) RegisterRoutes(handlers ...RouteHandler) {
	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	v1 := r.engine.Group("/api/v1")
	{
		for _, handler := range handlers {
			handler.RegisterRoutes(v1)
		}
	}
}
func (r *Router) RegisterProtectedRoutes(middleware gin.HandlerFunc, handlers ...RouteHandler) {
	if middleware == nil {
		return
	}
	v1 := r.engine.Group("/api/v1")
	protected := v1.Group("")
	protected.Use(middleware)
	{
		for _, handler := range handlers {
			handler.RegisterRoutes(protected)
		}
	}
}
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
func (r *Router) GetRouter() *Router {
	return r
}
