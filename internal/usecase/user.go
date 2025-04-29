package usecase

import (
	"clonecoding/internal/domain"
	"clonecoding/internal/dto"
	"clonecoding/internal/port"

	"github.com/google/uuid"
)

type UserUsecase struct {
	UserRepo port.UserRepository
	Hashing  port.Hashing
}

func (u *UserUsecase) CreateUser(userDto dto.CreateUserDTO) (*domain.User, error) {
	var user domain.User

	user.ID = uuid.New()

	user.Name = userDto.Name
	user.Email = userDto.Email

	password, err := u.Hashing.HashingPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	err = u.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.UserRepo.FindAllUser()
}

func (u *UserUsecase) GetUserById(userId uuid.UUID) (*domain.User, error) {
	return u.UserRepo.FindUserById(userId)
}
