package api

import (
	"fmt"
	"net/http"

	"github.com/Baroneee/EBVN_Assignment_1/internal/handler"
	"github.com/Baroneee/EBVN_Assignment_1/internal/service"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Start() error
	ServerHTTP(w http.ResponseWriter, r *http.Request)
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

func NewEngine(cfg *Config) Engine {
	e := &engine{
		app: gin.Default(),
		cfg: cfg,
	}
	e.initRoutes()
	return e
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%d", e.cfg.AppPort))
}

func (e *engine) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	e.app.ServeHTTP(w, r)
}

func (e *engine) initRoutes() {
    healthCheckSvc := service.NewHealthCheck(
        e.cfg.ServiceName,
        e.cfg.InstanceID,
    )

    healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)
    e.app.GET("/health", healthCheckHandler.Check)
}
