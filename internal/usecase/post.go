package usecase

import (
	"clonecoding/internal/domain"
	"clonecoding/internal/dto"
	"clonecoding/internal/port"

	"github.com/google/uuid"
)

type PostUsecase struct {
	PostRepo port.PostRepository
}

func (p *PostUsecase) CreatePost(createPostDTO *dto.CreatePostDTO, userId uuid.UUID, boardId uuid.UUID) (*domain.Post, error) {
	var post domain.Post
	post.Author = userId
	post.PostedBoard = boardId
	post.Title = createPostDTO.Title
	post.Body = createPostDTO.Body

	err := p.PostRepo.CreatePost(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostUsecase) FindAllPost() ([]*domain.Post, error) {
	boards, err := p.PostRepo.FindAllPost()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (p *PostUsecase) FindPostById(postId uuid.UUID) (*domain.Post, error) {
	board, err := p.PostRepo.FindPostById(postId)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func (p *PostUsecase) FindPostByBoardId(boardId uuid.UUID) ([]*domain.Post, error) {
	boards, err := p.PostRepo.FindPostByBoardId(boardId)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (p *PostUsecase) FindPostByUserId(userId uuid.UUID) ([]*domain.Post, error) {
	boards, err := p.PostRepo.FindPostByUserId(userId)
	if err != nil {
		return nil, err
	}

	return boards, nil
}
