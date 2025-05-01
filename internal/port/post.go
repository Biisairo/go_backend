package port

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
)

type PostRepository interface {
	CreatePost(post *domain.Post) error
	FindAllPost() ([]*domain.Post, error)
	FindPostById(postId uuid.UUID) (*domain.Post, error)
	FindPostByBoardId(boardId uuid.UUID) ([]*domain.Post, error)
	FindPostByUserId(userId uuid.UUID) ([]*domain.Post, error)
}
