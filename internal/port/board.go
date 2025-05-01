package port

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
)

type BoardRepository interface {
	CreateBoard(board *domain.Board) error
	FindAllBoard() ([]*domain.Board, error)
	FindBoardById(boardId uuid.UUID) (*domain.Board, error)
	FindUserByName(name string) ([]*domain.Board, error)
}
