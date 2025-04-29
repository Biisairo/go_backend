package database

import (
	"clonecoding/internal/domain"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func (r *AuthRepositoryImpl) CreateRefreshToken(refreshToken *domain.RefreshToken) error {
	return r.DB.Create(&refreshToken).Error
}

func (r *AuthRepositoryImpl) FindRefreshToken(refreshTokenStr string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.DB.First(&refreshToken, "token = ?", refreshTokenStr).Error
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *AuthRepositoryImpl) DeleteRefreshToken(refreshTokenStr string) error {
	err := r.DB.Delete(&domain.RefreshToken{}, "token = ?", refreshTokenStr).Error
	if err != nil {
		return err
	}

	return nil
}
