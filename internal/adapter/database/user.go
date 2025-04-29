package database

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepositoryImpl) FindAllUser() ([]*domain.User, error) {
	var users []*domain.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindUserById(userId uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
