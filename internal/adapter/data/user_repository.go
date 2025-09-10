package data

import (
	"errors"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
	"gorm.io/gorm"
)

type userRepository struct {
	config *config.DBConfig
	db     *gorm.DB
}

func NewUserRepository(c *config.Configuration, db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		config: c.DBConfig,
		db:     db,
	}
}

func (r *userRepository) Health(check chan<- *models.ComponentCheck) {
	// Implement the logic to check database connection health
	status := models.StatusPass
	if err := r.checkPing(); err != nil {
		status = models.StatusFail
	}

	check <- &models.ComponentCheck{
		Name:   r.config.Database,
		Type:   models.TypeDatastore,
		Status: status,
	}
}

func (r *userRepository) Save(user *models.User) error {
	// Implement the logic to save the user to the database
	return nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	// Implement the logic to find a user by ID from the database
	return &models.User{}, nil
}

func (r *userRepository) checkPing() error {
	if r.db == nil {
		return errors.New("database connection fail")
	}
	if pinger, ok := r.db.ConnPool.(interface{ Ping() error }); ok {
		return pinger.Ping()
	}

	return nil
}
