package interfaces

import "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"

type UserRepository interface {
	Health(chan<- *models.ComponentCheck)
	Save(user *models.User) error
	FindByID(id string) (*models.User, error)
}

type MessageProducer interface {
	Health(chan<- *models.ComponentCheck)
	Publish(topic string, message []byte) error
}
