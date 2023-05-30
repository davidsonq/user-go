package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var goblaToken string
var TestCount int

func TestLoginUser(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/api/users/login", handlers.LoginUserHandler)

	t.Run("SuccessLoginUser", func(t *testing.T) {
		requestBody, responseBody := mock.MockUserSucessLogin()

		jsonBody, _ := json.Marshal(requestBody)

		req, err := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.NotZero(t, len(responseBody.Token))

		goblaToken = responseBody.Token
		TestCount++
	})

	t.Run("TestingErrosValidation", func(t *testing.T) {

		testeCases := mock.MockUsersErrorValidationLogin()

		for _, testCase := range testeCases {

			jsonBody, _ := json.Marshal(testCase.RequestBody)

			req, err := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			assert.Contains(t, rec.Body.String(), "error")
			TestCount++
		}
	})

	t.Run("TestingErrosValidation", func(t *testing.T) {

		testeCases := mock.MockUserForbiden()

		for _, testCase := range testeCases {

			jsonBody, _ := json.Marshal(testCase.RequestBody)

			req, err := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusForbidden, rec.Code)

			assert.Contains(t, rec.Body.String(), "error")
			TestCount++
		}
	})

}
