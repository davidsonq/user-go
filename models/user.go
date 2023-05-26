package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid" binding:"required,uuid4"`
	Nickname  string         `json:"nickname" gorm:"size:50;unique" binding:"required,min=3,max=50"`
	Email     string         `json:"email" gorm:"size:50;unique" binding:"required,min=3,max=50,email"`
	Password  string         `json:"password" gorm:"password" binding:"required,min=6,max=16"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"deteled_at" gorm:"index"`
}
