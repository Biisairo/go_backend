package port

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	FindAllUser() ([]*domain.User, error)
	FindUserById(userId uuid.UUID) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
}
