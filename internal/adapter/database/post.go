package database

import (
	"clonecoding/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func (r *PostRepositoryImpl) CreatePost(post *domain.Post) error {
	return r.DB.Create(&post).Error
}

func (r *PostRepositoryImpl) FindAllPost() ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepositoryImpl) FindPostById(postId uuid.UUID) (*domain.Post, error) {
	var post domain.Post
	err := r.DB.First(&post, "id = ?", postId).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepositoryImpl) FindPostByBoardId(boardId uuid.UUID) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.DB.Find(&posts, "PostedBoard = ?", boardId).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepositoryImpl) FindPostByUserId(userId uuid.UUID) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.DB.Find(&posts, "Author = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
