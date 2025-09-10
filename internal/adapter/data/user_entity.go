package data

import "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"

type UserEntity struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func toUserEntity(user *models.User) *UserEntity {
	if user == nil {
		return nil
	}

	return &UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func fromUserEntity(entity *UserEntity) *models.User {
	if entity == nil {
		return nil
	}

	return &models.User{
		ID:    entity.ID,
		Name:  entity.Name,
		Email: entity.Email,
	}
}
