package http

import (
	"clonecoding/internal/dto"
	"clonecoding/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	token, refreshToken, err := a.AuthUseCase.Login(&req)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, map[string]any{"access_token": token, "refresh_token": refreshToken})
}

func (a *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	token, refreshToken, err := a.AuthUseCase.Refresh(&req)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, map[string]any{"access_token": token, "refresh_token": refreshToken})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.AuthUseCase.Logout((&req)); err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, nil)
}

// func (a *AuthHandler) Logout(c *gin.Context) {
// 	var req dto.RefreshRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		Fail(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	err := a.AuthUseCase.Logout(req)
// 	if err != nil {
// 		Fail(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	Success(c, nil)
// }
