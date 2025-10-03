package user

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

type UserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func toUserDTO(user *models.User) *UserDTO {
	if user == nil {
		return nil
	}

	return &UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func fromUserDTO(dto *UserDTO) *models.User {
	if dto == nil {
		return nil
	}

	return &models.User{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
	}
}
