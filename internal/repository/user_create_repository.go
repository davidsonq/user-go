package repository

import (
	"log"

	"github.com/davidsonq/user-go/internal/db"
	"github.com/davidsonq/user-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(u *models.User) (*models.UserResponse, error) {
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

	return &models.UserResponse{
		ID:        u.ID,
		Nickname:  u.Nickname,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil

}
