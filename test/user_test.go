package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"user-go/handlers"
	"user-go/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatedUser(t *testing.T) {
	r := gin.Default()
	r.POST("/api/users", handlers.CreateUser)
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
		assert.Empty(t, responseBody.DeleteAt)
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

		}
	})
}
