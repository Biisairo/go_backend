package http

import (
	"clonecoding/internal/dto"
	"clonecoding/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BoardHandler struct {
	BoardUseCase *usecase.BoardUsecase
}

func (b *BoardHandler) CreateBoard(c *gin.Context) {
	var boardDto dto.CreateBoardDTO
	err := c.ShouldBindJSON(&boardDto)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
	}

	user, err := b.BoardUseCase.CreateBoard(&boardDto)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
	}

	Success(c, user)
}

func (b *BoardHandler) GetAllBoard(c *gin.Context) {
	boards, err := b.BoardUseCase.FindAllBoard()
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
	}

	Success(c, boards)
}

func (b *BoardHandler) GetBoard(c *gin.Context) {
	id := c.Param("id")

	boardId, err := uuid.Parse(id)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
	}

	board, err := b.BoardUseCase.FindBoardById(boardId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
	}

	Success(c, board)
}
