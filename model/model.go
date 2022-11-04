package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64         `gorm:"primaryKey" json:"id" `
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

var Models = []any{
	&FileMetadata{},
}
