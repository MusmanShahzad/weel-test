package app

import (
	"weel-backend/config"
	"weel-backend/internal/container"
	"weel-backend/internal/database"
	"weel-backend/internal/module/auth"
	"weel-backend/internal/module/feature_flag"
	"weel-backend/internal/module/order"
	"weel-backend/internal/module/user"
)

type App struct {
	container *container.Container
	config    *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {
	app := &App{
		config:    cfg,
		container: container.NewContainer(),
	}
	app.container.DB = database.DB
	app.registerModules()
	if err := app.container.Initialize(); err != nil {
		return nil, err
	}
	return app, nil
}
func (a *App) registerModules() {
	a.container.RegisterModule(feature_flag.NewFeatureFlagModule())
	a.container.RegisterModule(auth.NewAuthModule())
	a.container.RegisterModule(order.NewOrderModule(a.config))
	a.container.RegisterModule(user.NewUserModule())
}
func (a *App) GetRouter() *container.Container {
	return a.container
}
func (a *App) Close() error {
	return database.Close()
}
