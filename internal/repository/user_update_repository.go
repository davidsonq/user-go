package repository

import (
	"errors"
	"log"

	"github.com/davidsonq/user-go/internal/db"
	"github.com/davidsonq/user-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(id string, u *models.UpdateUser) (*models.UserResponse, error) {
	if u.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.DefaultCost)

		if err != nil {
			log.Fatal(err)
		}

		hashedPasswordStr := string(hashedPassword)
		u.Password = &hashedPasswordStr
	}

	db := db.ConnectionDB()
	var user models.User
	var existUser models.User

	err := db.Where("nickname = ? AND id != ?", u.Nickname, id).First(&existUser).Error

	if err == nil {
		return nil, errors.New("The nickname is already in use")
	}

	err = db.Where("email = ? AND id != ?", u.Email, id).First(&existUser).Error
	if err == nil {
		return nil, errors.New("The email is already in use")
	}

	if err := db.Table("users").First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}

	db.Model(&user).Updates(u)

	return &models.UserResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
