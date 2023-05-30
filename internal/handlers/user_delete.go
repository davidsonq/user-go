package handlers

import (
	"net/http"

	"github.com/davidsonq/user-go/internal/repository"
	"github.com/gin-gonic/gin"
)

// @Summary Delete User
// @Description Deleted User
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} models.ErrosNoBody "User not exist"
// @Failure 401 {object} models.ErrosNoBody "When you pass an invalid token in the request header or it was not sent for authentication."
// @Router /users/:id [delete]
func DeleteUserHandler(c *gin.Context) {

	userID := c.GetString("userID")

	if userID != c.Param("id") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access."})
		return
	}
	err := repository.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deleting users is not allowed!"})
		return
	}

	c.Status(http.StatusNoContent)

}
