package handler

import (
	"net/http"

	"github.com/Baroneee/EBVN_Assignment_1/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthCheck interface {
	Check(c *gin.Context)
}

type healthCheck struct {
	svc service.HealthCheck
}

func NewHealthCheck(svc service.HealthCheck) HealthCheck {
	return &healthCheck{
		svc: svc,
	}
}

func (h *healthCheck) Check(c *gin.Context) {
	result := h.svc.Check()
	c.JSON(http.StatusOK, result)
}