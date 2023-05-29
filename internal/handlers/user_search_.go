package handlers

import (
	"net/http"

	"github.com/davidsonq/user-go/internal/repository"
	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Seach profile user
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.UserResponse
// @Failure 401 {object} models.ErrosNoBody "When you pass an invalid token in the request header or it was not sent for authentication."
// @Router /users/profile [get]
func GetProfileUserHandler(c *gin.Context) {
	userID := c.GetString("userID")

	u, err := repository.GetProfile(&userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action!"})
		return
	}

	c.JSON(http.StatusOK, u)

}
