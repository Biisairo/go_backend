package dto

import "github.com/google/uuid"

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`

	Password string `json:"password" binding:"required"`
}

type UserUriDTO struct {
	UserId uuid.UUID `uri:"board_id" binding:"required"`
}
