package handlers

import (
	"net/http"

	"github.com/davidsonq/user-go/internal/models"
	"github.com/davidsonq/user-go/internal/repository"
	"github.com/davidsonq/user-go/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Session User
// @Description Create new session
// @Tags session
// @Accept json
// @Produce json
// @Param request body models.LoginUser true "Request body"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrosNoBody "The request body is empty or missing Email, Password in the request body, the email has to be a valid email and the password has at least 6 characters."
// @Failure 403 {object} models.ErrosNoBody "This error is generated when trying to login with an invalid email or password."
// @Router /api/users/login [post]
func LoginUserHandler(c *gin.Context) {
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The request body is empty"})
		c.Abort()
		return
	}

	var session models.LoginUser

	if err := c.ShouldBindJSON(&session); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validations.GetCustomErrorMessageUser(errs)})
		return
	}

	l, err := repository.LoginUser(&session)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email or Password invalid!"})
		return
	}

	c.JSON(http.StatusOK, l)

}
