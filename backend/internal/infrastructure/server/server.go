package server

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/routes"
	"go.uber.org/dig"
)

// CustomValidator wraps the go-playground validator for use with Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the struct using the go-playground validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewCustomValidator creates a new CustomValidator instance
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// EchoServer struct manages the Echo server configuration
type EchoServer struct {
	Route     *echo.Echo
	AppConfig *config.AppConfig
}

// Start starts the Echo server on the configured port
func (s *EchoServer) Start() error {
	return s.Route.Start(fmt.Sprintf("0.0.0.0:%s", s.AppConfig.AppPort))
}

// NewEchoServer initializes a new EchoServer with the given configuration and DI container
func NewEchoServer(c *config.AppConfig, diContainer *dig.Container) *EchoServer {
	e := echo.New()

	// Set the custom validator
	e.Validator = NewCustomValidator()

	// Use the routes package to configure routes
	routes.RegisterRoutes(e, diContainer, c)

	return &EchoServer{
		Route:     e,
		AppConfig: c,
	}
}
