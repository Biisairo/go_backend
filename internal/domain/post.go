package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Author uuid.UUID `gorm:"index" binding:"required"`
	User   User      `gorm:"foreignKey:Author;references:ID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`

	PostedBoard uuid.UUID `gorm:"index" binding:"required"`
	Board       Board     `gorm:"foreignKey:PostedBoard;references:ID;constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`

	Title string `binding:"required"`
	Body  string
}
