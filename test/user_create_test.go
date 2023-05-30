package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"testing"

	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/internal/models"
	"github.com/davidsonq/user-go/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatedUser(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.POST("/api/users", handlers.CreateUserHandle)
	t.Run("SuccessfulCreation", func(t *testing.T) {

		requestBody, responseBody := mock.MockUserSucess()

		jsonBody, _ := json.Marshal(requestBody)

		req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		_, err = uuid.Parse(responseBody.ID)
		assert.NoError(t, err)
		assert.Equal(t, requestBody.Email, responseBody.Email)
		assert.Equal(t, requestBody.Nickname, responseBody.Nickname)
		assert.Empty(t, responseBody.Password)
		assert.Equal(t, reflect.TypeOf(requestBody.CreatedAt), reflect.TypeOf(responseBody.CreatedAt))
		assert.Equal(t, reflect.TypeOf(requestBody.UpdatedAt), reflect.TypeOf(responseBody.UpdatedAt))

		TestCount++
	})

	t.Run("TestingErrosValidation", func(t *testing.T) {

		testeCases := mock.MockUsersErrorValidation()

		for _, testCase := range testeCases {

			jsonBody, _ := json.Marshal(testCase.RequestBody)

			req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			assert.Contains(t, rec.Body.String(), "error")
			TestCount++
		}

	})

	t.Run("TestingErrosDuplication", func(t *testing.T) {
		testCases := mock.MockUserDuplicate()

		for _, testCase := range testCases {

			jsonBody, _ := json.Marshal(testCase.RequestBody)

			req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusConflict, rec.Code)

			assert.Contains(t, rec.Body.String(), "error")
			TestCount++
		}
	})
}

func BenchmarkCreateUserHandle(b *testing.B) {

	r := gin.Default()
	r.POST("/api/users", handlers.CreateUserHandle)
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.Run("Succes", func(b *testing.B) {
		for n := 1; n < b.N; n++ {
			requestBody := models.User{
				ID:       uuid.New().String(),
				Nickname: fmt.Sprintf("test%v", n),
				Email:    fmt.Sprintf("test%v@example.com", n),
				Password: "123456",
			}
			jsonBody, _ := json.Marshal(requestBody)
			req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)
			assert.Equal(b, http.StatusCreated, rec.Code)

		}
	})
	b.Run("Conflit", func(b *testing.B) {

		for n := 1; n < b.N; n++ {
			requestBody := models.User{
				ID:       uuid.New().String(),
				Nickname: fmt.Sprintf("test%v", n),
				Email:    fmt.Sprintf("test%v@example.com", n),
				Password: "123456",
			}
			jsonBody, _ := json.Marshal(requestBody)
			req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)
			assert.Equal(b, http.StatusConflict, rec.Code)
		}
	})

}
