package usecase

import (
	"clonecoding/internal/domain"
	"clonecoding/internal/dto"
	"clonecoding/internal/port"

	"github.com/google/uuid"
)

type BoardUsecase struct {
	BoardRepo port.BoardRepository
}

func (b *BoardUsecase) CreateBoard(createBoardDTO *dto.CreateBoardDTO) (*domain.Board, error) {
	var board domain.Board
	board.ID = uuid.New()
	board.Name = createBoardDTO.Name

	err := b.BoardRepo.CreateBoard(&board)
	if err != nil {
		return nil, err
	}

	return &board, nil
}

func (b *BoardUsecase) FindAllBoard() ([]*domain.Board, error) {
	boards, err := b.BoardRepo.FindAllBoard()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (b *BoardUsecase) FindBoardById(boardId uuid.UUID) (*domain.Board, error) {
	board, err := b.BoardRepo.FindBoardById(boardId)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func (b *BoardUsecase) FindUserByName(name string) ([]*domain.Board, error) {
	boards, err := b.BoardRepo.FindUserByName(name)
	if err != nil {
		return nil, err
	}

	return boards, nil
}
