package routes
import (
	"weel-backend/internal/handler"
	"weel-backend/internal/router"
)
func SetupRoutes(
	userHandler *handler.UserHandler,
) *router.Router {
	r := router.NewRouter()
	r.RegisterRoutes(
		userHandler,
	)
	return r
}
