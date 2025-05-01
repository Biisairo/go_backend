package dto

import "github.com/google/uuid"

type CreatePostDTO struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

type PostUriDTO struct {
	PostId uuid.UUID `uri:"post_id" binding:"required"`
}
