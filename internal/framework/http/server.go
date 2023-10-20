// Package http implements HTTP server
// nolint: wrapcheck
package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mirshahriar/marketplace/config"
	"github.com/mirshahriar/marketplace/internal/framework/http/middlewares"
	"github.com/mirshahriar/marketplace/internal/ports"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Adapter implements HTTP interface
type Adapter struct {
	config config.AppConfig
	// Echo is the HTTP server.
	echo *echo.Echo
	// Binder is the default binder for the HTTP server.
	binder echo.DefaultBinder
	// APIPort is the port to the application's business logic.
	api ports.APIPort
}

// NewAdapter creates a new Adapter struct and returns a pointer to it.
func NewAdapter(config config.AppConfig, api ports.APIPort) *Adapter {
	return &Adapter{
		config: config,
		echo:   echo.New(),
		api:    api,
	}
}

func (a Adapter) Run() {
	middlewares.Init(a.echo)

	a.registerAPI()

	a.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
}

func (a Adapter) registerAPI() {
	// Registering the routes those are accessible without authentication.
	unAuthGroup := a.echo.Group("")
	unAuthGroup.GET("/products", a.ListProduct)
	unAuthGroup.GET("/products/:product", a.GetProduct)

	// Registering the routes those are accessible with authentication.
	authGroup := a.echo.Group("", middlewares.AuthorizeUser(a.api))
	authGroup.POST("/products", a.CreateProduct)
	authGroup.PUT("/products/:product", a.UpdateProduct)
	authGroup.DELETE("/products/:product", a.DeleteProduct)
}
