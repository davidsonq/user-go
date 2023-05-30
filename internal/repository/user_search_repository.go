package repository

import (
	"github.com/davidsonq/user-go/internal/db"
	"github.com/davidsonq/user-go/internal/models"
)

func GetProfile(id *string) (*models.UserResponse, error) {

	db := db.ConnectionDB()
	var user models.UserResponse

	if err := db.Table("users").First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
