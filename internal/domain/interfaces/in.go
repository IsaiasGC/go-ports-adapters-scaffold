package interfaces

import "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"

type HealthService interface {
	Check() *models.HealthCheck
	CheckDependencies() *models.HealthCheck
}
type UserService interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
}
