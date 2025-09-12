package interfaces

import (
	"context"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

type HealthService interface {
	Check() *models.HealthCheck
	CheckDependencies(ctx context.Context) *models.HealthCheck
}
type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
}
