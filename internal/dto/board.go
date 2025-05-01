package dto

type CreateBoardDTO struct {
	Name string `json:"name" binding:"required"`
}
