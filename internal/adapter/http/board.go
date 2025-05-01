package http

import (
	"clonecoding/internal/dto"
	"clonecoding/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	BoardUseCase *usecase.BoardUsecase
}

func (b *BoardHandler) CreateBoard(c *gin.Context) {
	var boardDto dto.CreateBoardDTO
	err := c.ShouldBindJSON(&boardDto)
	if err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := b.BoardUseCase.CreateBoard(&boardDto)
	if err != nil {
		Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	Success(c, user)
}

func (b *BoardHandler) GetAllBoard(c *gin.Context) {
	boards, err := b.BoardUseCase.FindAllBoard()
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, boards)
}

func (b *BoardHandler) GetBoard(c *gin.Context) {
	var uri dto.BoardUriDTO
	if err := c.ShouldBindUri(&uri); err != nil {
		Fail(c, http.StatusBadRequest, err.Error())
		return
	}

	board, err := b.BoardUseCase.FindBoardById(uri.BoardId)
	if err != nil {
		Fail(c, http.StatusNotFound, err.Error())
		return
	}

	Success(c, board)
}
