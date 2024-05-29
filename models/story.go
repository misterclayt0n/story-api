package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Story struct {
	ID          string `gorm:"type:uuid;primary_key"`
	Title       string
	Description string
	Category    string
}

func (s *Story) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
