package api

import (
	"fmt"
	"net/http"

	"github.com/Baroneee/EBVN_Assignment_1/internal/handler"
	"github.com/Baroneee/EBVN_Assignment_1/internal/service"
	"github.com/gin-gonic/gin"
)

// Engine starts the HTTP server and handles HTTP requests.
type Engine interface {
	Start() error
	ServerHTTP(w http.ResponseWriter, r *http.Request)
}

// engine is the Gin-based implementation of Engine.
type engine struct {
	app *gin.Engine
	cfg *Config
}

// NewEngine creates an HTTP engine configured with cfg.
func NewEngine(cfg *Config) Engine {
	e := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	e.initRoutes()
	return e
}

// Start runs the HTTP server on the configured application port.
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%d", e.cfg.AppPort))
}

// ServerHTTP serves an HTTP request through the configured Gin engine.
func (e *engine) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}

// initRoutes registers the application's HTTP routes.
func (e *engine) initRoutes() {
	healthCheckSvc := service.NewHealthCheck(
		e.cfg.ServiceName,
		e.cfg.InstanceID,
	)

	healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)
	e.app.GET("/health", healthCheckHandler.Check)
}
