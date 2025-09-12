package application

import (
	"context"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

const TOTAL_DEPENDENCIES = 2

type healthService struct {
	info     *config.ProjectInfo
	repo     interfaces.UserRepository
	producer interfaces.MessageProducer
}

func NewHealthService(c *config.Configuration,
	r interfaces.UserRepository, p interfaces.MessageProducer) interfaces.HealthService {
	return &healthService{
		info:     c.ProjectInfo,
		repo:     r,
		producer: p,
	}
}

func (s *healthService) Check() *models.HealthCheck {
	return &models.HealthCheck{
		Version: s.info.Version,
		Status:  models.StatusPass,
	}
}

func (s *healthService) CheckDependencies(ctx context.Context) *models.HealthCheck {
	dependencies := s.getHealthDependencies(ctx)

	return &models.HealthCheck{
		Version: s.info.Version,
		Status:  s.getCheckStatus(dependencies),
		Checks:  dependencies,
	}
}

func (s *healthService) getHealthDependencies(ctx context.Context) []*models.ComponentCheck {
	dependencies := make([]*models.ComponentCheck, 0, TOTAL_DEPENDENCIES)

	checks := make(chan *models.ComponentCheck, TOTAL_DEPENDENCIES)
	defer close(checks)

	go s.repo.HealthCheck(ctx, checks)
	go s.producer.HealthCheck(ctx, checks)

	for len(dependencies) < TOTAL_DEPENDENCIES {
		if check := <-checks; check != nil {
			dependencies = append(dependencies, check)
		}
	}

	return dependencies
}

func (s *healthService) getCheckStatus(checks []*models.ComponentCheck) models.HealthStatus {
	status := models.StatusPass

	for _, check := range checks {
		if check.IsFail() {
			if check.Type.IsFatal() {
				status = models.StatusFail
			} else if status != models.StatusFail {
				status = models.StatusWarn
			}
		}
	}

	return status
}
