package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid" validate:"required,uuid4"`
	NickName  string         `json:"nickname" gorm:"size:50;unique" validate:"required,min=3,max=50"`
	Email     string         `json:"email" gorm:"size:50;unique" validate:"required,min=3,max=50"`
	Password  string         `json:"" gorm:"password;size:16" validate:"required,min=6,max=16"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"" gorm:"index"`
}
