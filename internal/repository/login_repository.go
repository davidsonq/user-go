package repository

import (
	"os"
	"time"

	"github.com/davidsonq/user-go/internal/db"
	"github.com/davidsonq/user-go/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Jwt struct {
	Key []byte
	T   *jwt.Token
}

func LoginUser(l *models.LoginUser) (*models.LoginResponse, error) {
	db := db.ConnectionDB()
	var user models.User

	if err := db.Table("users").First(&user, "email=?", l.Email).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))

	if err != nil {
		return nil, err
	}

	jwtNew := Jwt{
		Key: []byte(os.Getenv("SECRET_KEY")),
		T: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "Auth-user-go",
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}),
	}

	tokenString, err := jwtNew.T.SignedString(jwtNew.Key)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: tokenString,
	}, nil
}
