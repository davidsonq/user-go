package handlers

import (
	"net/http"

	"github.com/davidsonq/user-go/models"
	"github.com/davidsonq/user-go/repository"
	"github.com/davidsonq/user-go/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// @Summary Create user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.ExampleInputUser true "Request body"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.ErrosNoBody "The body of the request is empty or lack of Nickname, Email, Password in the body of the request, the email has to be a valid email and nickname has to have at least 3 characters and password at least 6 "
// @Failure 401 {object} models.ErrosNoBody "This error is generated when trying to register a nickname or email already registered"
// @Router /api/users [post]
func CreateUserHandle(c *gin.Context) {
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
