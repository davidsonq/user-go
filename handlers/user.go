package handlers

import (
	"net/http"
	"user-go/models"
	"user-go/repository"
	"user-go/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request body is empty"})
		c.Abort()
		return
	}

	var user models.User
	user.ID = uuid.New().String()

	if err := c.ShouldBindJSON(&user); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validations.GetCustomErrorMessageUser(errs)})
		return
	}

	u, err := repository.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": validations.DuplicateErrorUser(err)})
		return
	}

	c.JSON(http.StatusCreated, u)
}
