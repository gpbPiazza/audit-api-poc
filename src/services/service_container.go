package services

import (
	"github.com/gpbPiazza/src/interfaces"
	"github.com/gpbPiazza/src/services/health_check"
)

type ServiceContainer struct {
	healthCheckService interfaces.HealthCheckService
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		healthCheckService: health_check.NewHealthCheckService(),
	}
}
