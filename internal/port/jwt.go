package port

import (
	"clonecoding/internal/adapter/jwt"

	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	ParseToken(tokenStr string) (*jwt.Claims, error)
}
