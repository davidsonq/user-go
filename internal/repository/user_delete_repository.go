package repository

import (
	"github.com/davidsonq/user-go/internal/db"
	"github.com/davidsonq/user-go/internal/models"
)

func DeleteUser(id string) error {
	db := db.ConnectionDB()

	var user models.UserResponse

	if err := db.Table("users").First(&user, "id=? AND deleted_at IS NULL", id).Delete(&user, "id=?", id).Error; err != nil {
		return err
	}

	return nil
}
