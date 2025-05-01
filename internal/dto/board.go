package dto

import "github.com/google/uuid"

type CreateBoardDTO struct {
	Name string `json:"name" binding:"required"`
}

type BoardUriDTO struct {
	BoardId uuid.UUID `uri:"board_id" binding:"required"`
}
