package database

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardRepositoryImpl struct {
	DB *gorm.DB
}

func (r *BoardRepositoryImpl) CreateBoard(board *domain.Board) error {
	return r.DB.Create(&board).Error
}

func (r *BoardRepositoryImpl) FindAllBoard() ([]*domain.Board, error) {
	var boards []*domain.Board
	err := r.DB.Find(&boards).Error
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *BoardRepositoryImpl) FindBoardById(boardId uuid.UUID) (*domain.Board, error) {
	var board domain.Board
	err := r.DB.First(&board, "id = ?", boardId).Error
	if err != nil {
		return nil, err
	}

	return &board, nil
}

func (r *BoardRepositoryImpl) FindUserByName(name string) ([]*domain.Board, error) {
	var boards []*domain.Board
	err := r.DB.Find(&boards, "name = ?", name).Error
	if err != nil {
		return nil, err
	}

	return boards, nil
}
