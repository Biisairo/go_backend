package http

import (
	"clonecoding/internal/dto"
	"clonecoding/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserUseCase *usecase.UserUsecase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDto dto.CreateUserDTO
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.UserUseCase.CreateUser(userDto)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
	}

	Success(c, user)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.UserUseCase.GetAllUsers()
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
	}

	Success(c, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.UserUseCase.GetUserById(userId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
	}

	Success(c, user)
}
