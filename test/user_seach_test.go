package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/internal/middlewares"
	"github.com/davidsonq/user-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/api/users/profile", middlewares.AuthMiddleware(), handlers.GetProfileUser)

	t.Run("SucessGetProfile", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/users/profile", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		requestBody := "test@example.com"
		var responseBody models.User

		err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, requestBody, responseBody.Email)
		assert.Empty(t, responseBody.Password)

	})

	t.Run("ErrorNoAuthorization", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/users/profile", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
	})

	t.Run("ErrorNoTokenInvalid", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/users/profile", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "erroToken"))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
	})
}
