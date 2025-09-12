package application

import (
	"context"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/validations"
)

type userService struct {
	repo     interfaces.UserRepository
	producer interfaces.MessageProducer
}

func NewUserService(r interfaces.UserRepository, p interfaces.MessageProducer) interfaces.UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	if err := validations.ValidateCreateUser(user); err != nil {
		return err
	}
	if err := s.repo.Save(ctx, user); err != nil {
		return err
	}
	err := s.producer.Publish(ctx, "user.created", []byte(user.ID))

	return err
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}
