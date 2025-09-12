package data

import (
	"context"
	"errors"
	"time"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/apperror"
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

func (r *userRepository) HealthCheck(ctx context.Context, check chan<- *models.ComponentCheck) {
	st := time.Now()
	health := &models.ComponentCheck{
		Name:   r.config.Database,
		Type:   models.TypeDatastore,
		Status: models.StatusPass,
	}

	if err := r.checkPing(ctx); err != nil {
		health.Status = models.StatusFail
		health.Output = err.Error()
	}

	health.Time = time.Since(st)
	check <- health
}

func (r *userRepository) Save(ctx context.Context, model *models.User) error {
	tx := r.db.WithContext(ctx)

	user := toUserEntity(model)
	err := tx.Create(user).Error
	if err != nil {
		return apperror.NewError(apperror.CodeInternalError,
			"error saving user",
			err)
	}

	model.ID = user.ID

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	tx := r.db.WithContext(ctx)

	user := &UserEntity{}

	err := tx.Where("id = ?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NewError(apperror.CodeNotFound,
				"user not found",
				err)
		}
		return nil, apperror.NewError(apperror.CodeInternalError,
			"error fetching user",
			err)
	}

	return fromUserEntity(user), nil
}

func (r *userRepository) checkPing(ctx context.Context) error {
	if r.db == nil {
		return errors.New("database connection fail")
	}

	tx := r.db.WithContext(ctx)

	if pinger, ok := tx.ConnPool.(interface{ Ping() error }); ok {
		return pinger.Ping()
	}

	return nil
}
