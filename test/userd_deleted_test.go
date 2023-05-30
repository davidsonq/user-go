package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/internal/middlewares"
	"github.com/davidsonq/user-go/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.DELETE("/api/users/:id", middlewares.AuthMiddleware(), handlers.DeleteUserHandler)

	t.Run("ErrorNoAuthorization", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v", profileId), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

	t.Run("ErrorNoTokenInvalid", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v", profileId), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "erroToken"))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

	t.Run("ErrorIdInvalid", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v", "idInvalid"), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "erroToken"))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

	t.Run("SucessDeleteUser", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v", profileId), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		TestCount++
	})
	t.Run("ErrorUserDelete", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/users/%v", profileId), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		TestCount++
	})
}

func TestGetProfileUserDeleted(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/api/users/profile", middlewares.AuthMiddleware(), handlers.GetProfileUserHandler)

	req, err := http.NewRequest("GET", "/api/users/profile", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
	TestCount++
}

func TestTryLoginUserDeleted(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/api/users/login", handlers.LoginUserHandler)

	requestBody, _ := mock.MockUserSucessLogin()

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
	TestCount++
}

func TestTryUpdateUserDeleted(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.PATCH("/api/users/:id", middlewares.AuthMiddleware(), handlers.UpdatedUserHandle)

	updateCases := mock.MockUpdateSucess()[0]

	jsonBody, _ := json.Marshal(updateCases.RequestBody)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusConflict, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
	TestCount++

}

func TestMain(m *testing.M) {
	m.Run()
	fmt.Printf("Quantidade de test aprovados foram %v", TestCount)
}
