package repository

import (
	"log"
	"time"
	"user-go/db"
	"user-go/models"

	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(u *models.User) (*UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hashedPassword)

	db := db.ConnectionDB()

	if err := db.Create(u).Error; err != nil {
		return nil, err
	}
	u.Password = ""

	return &UserResponse{
		ID:        u.ID,
		Nickname:  u.Nickname,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil

}
