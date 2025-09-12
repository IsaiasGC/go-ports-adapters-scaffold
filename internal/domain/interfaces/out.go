package interfaces

import (
	"context"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

type UserRepository interface {
	HealthCheck(ctx context.Context, check chan<- *models.ComponentCheck)
	Save(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
}

type MessageProducer interface {
	HealthCheck(ctx context.Context, check chan<- *models.ComponentCheck)
	Publish(ctx context.Context, topic string, message []byte) error
}
