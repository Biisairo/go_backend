package dto

type CreatePostDTO struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}
