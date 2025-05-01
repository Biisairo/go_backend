package domain

import "github.com/google/uuid"

type Board struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Name string `gorm:"index"`
}
