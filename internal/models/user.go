package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid" binding:"required,uuid4"`
	Nickname  string    `json:"nickname" gorm:"size:50;unique" binding:"required,min=3,max=50"`
	Email     string    `json:"email" gorm:"size:50;unique" binding:"required,max=50,email"`
	Password  string    `json:"password" gorm:"password" binding:"required,min=6,max=16"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExampleInputUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ErrosNoBody struct {
	Error string `json:"error"`
}

type UpdateUser struct {
	Nickname  *string   `json:"nickname" binding:"omitempty,min=3,max=50"`
	Email     *string   `json:"email" binding:"omitempty,max=50,email"`
	Password  *string   `json:"password" binding:"omitempty,min=6,max=16"`
	UpdatedAt time.Time `json:"updated_at"`
}
