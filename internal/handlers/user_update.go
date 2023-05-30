package handlers

import (
	"net/http"

	"github.com/davidsonq/user-go/internal/models"
	"github.com/davidsonq/user-go/internal/repository"
	"github.com/davidsonq/user-go/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Update User
// @Description Update info user
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ErrosNoBody "The body of the request is empty or lack of Nickname, Email, Password in the body of the request, the email has to be a valid email and nickname has to have at least 3 characters and password at least 6."
// @Failure 409 {object} models.ErrosNoBody "User not exist"
// @Failure 401 {object} models.ErrosNoBody "When you pass an invalid token in the request header or it was not sent for authentication."
// @Router /users/:id [patch]
func UpdatedUserHandle(c *gin.Context) {
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request body is empty"})
		c.Abort()
		return
	}

	userID := c.GetString("userID")

	if userID != c.Param("id") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access."})
		return
	}
	var updateUser models.UpdateUser

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validations.GetCustomErrorMessageUser(errs)})
		return
	}
	u, err := repository.UpdateUser(c.Param("id"), &updateUser)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": validations.DuplicateErrorUser(err)})
		return
	}

	c.JSON(http.StatusOK, u)

}
