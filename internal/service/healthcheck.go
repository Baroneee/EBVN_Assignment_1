package service

import (
	"uuid"
)

//go:generate mockery --name=HealthCheck --filename=healthcheck.go --outpkg healthcheck_svc_mocks
type HealthResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

type HealthCheck interface {
	Check() HealthResponse
}

type healthCheck struct {
	serviceName string
	instanceID  string
}

func NewHealthCheck(serviceName, instanceID string) HealthCheck {
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	return &healthCheck{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (h *healthCheck) Check() HealthResponse {
	return HealthResponse{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceID:  h.instanceID,
	}
}
