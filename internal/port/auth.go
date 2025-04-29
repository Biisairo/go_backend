package port

import (
	"clonecoding/internal/domain"
)

type AuthRepository interface {
	CreateRefreshToken(refreshToken *domain.RefreshToken) error
	FindRefreshToken(refreshTokenStr string) (*domain.RefreshToken, error)
	DeleteRefreshToken(refreshTokenStr string) error
}
