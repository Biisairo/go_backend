package usecase

import (
	"clonecoding/internal/domain"
	"clonecoding/internal/dto"
	"clonecoding/internal/port"
	"errors"
	"time"
)

type AuthUseCase struct {
	UserRepo   port.UserRepository
	AuthRepo   port.AuthRepository
	JWTService port.JWTService
	Hashing    port.Hashing
}

func (a *AuthUseCase) Login(loginReq *dto.LoginRequest) (string, string, error) {
	user, err := a.UserRepo.FindUserByEmail(loginReq.Email)
	if err != nil {
		return "", "", err
	}

	if !a.Hashing.CheckPasswordHash(loginReq.Password, user.Password) {
		return "", "", errors.New("error: password not matched")
	}

	token, err := a.JWTService.GenerateToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := a.JWTService.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshTokenDomain := domain.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = a.AuthRepo.CreateRefreshToken(&refreshTokenDomain)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, err
}

func (a *AuthUseCase) Refresh(refreshReq *dto.RefreshRequest) (string, string, error) {
	refreshTokenStr := refreshReq.RefreshToken
	claims, err := a.JWTService.ParseToken(refreshTokenStr)
	if err != nil {
		return "", "", err
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("expired refresh token")
	}

	token, err := a.JWTService.GenerateToken(claims.UserID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := a.JWTService.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return "", "", err
	}

	refreshTokenDomain := domain.RefreshToken{
		UserID:    claims.UserID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err = a.AuthRepo.CreateRefreshToken(&refreshTokenDomain)
	if err != nil {
		return "", "", err
	}

	err = a.AuthRepo.DeleteRefreshToken(refreshTokenStr)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (a *AuthUseCase) Logout(refreshReq *dto.RefreshRequest) error {
	return a.AuthRepo.DeleteRefreshToken(refreshReq.RefreshToken)
}
