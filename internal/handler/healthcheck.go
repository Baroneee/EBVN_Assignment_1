package handler

import (
	"net/http"

	"github.com/Baroneee/EBVN_Assignment_1/internal/service"
	"github.com/gin-gonic/gin"
)

// HealthCheck handles health check HTTP requests.
type HealthCheck interface {
	Check(c *gin.Context)
}

// healthCheck is the HTTP handler for health check requests.
type healthCheck struct {
	svc service.HealthCheck
}

// NewHealthCheck creates a health check handler backed by svc.
func NewHealthCheck(svc service.HealthCheck) HealthCheck {
	return &healthCheck{
		svc: svc,
	}
}

// Check writes the service health status as a JSON response.
func (h *healthCheck) Check(c *gin.Context) {
	result := h.svc.Check()
	c.JSON(http.StatusOK, result)
}
