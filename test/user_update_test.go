package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/internal/middlewares"
	"github.com/davidsonq/user-go/internal/models"
	"github.com/davidsonq/user-go/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.PATCH("/api/users/:id", middlewares.AuthMiddleware(), handlers.UpdatedUserHandle)

	t.Run("SucessUpdateUser", func(t *testing.T) {
		updateCases := mock.MockUpdateSucess()
		for _, updateCase := range updateCases {
			jsonBody, _ := json.Marshal(updateCase.RequestBody)

			req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)

			var responseBody models.User
			err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			_, err = uuid.Parse(responseBody.ID)
			assert.NoError(t, err)
			assert.Equal(t, *updateCase.RequestBody.Email, responseBody.Email)
			assert.Equal(t, *updateCase.RequestBody.Nickname, responseBody.Nickname)
			assert.Empty(t, responseBody.Password)
			assert.Equal(t, reflect.TypeOf(updateCase.RequestBody.UpdatedAt), reflect.TypeOf(responseBody.UpdatedAt))
			TestCount++
		}

	})

	t.Run("ErrorNoTokenInvalid", func(t *testing.T) {
		updateCases := mock.MockUpdateSucess()[len(mock.MockUpdateSucess())-1]

		jsonBody, _ := json.Marshal(updateCases.RequestBody)

		req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
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
		updateCases := mock.MockUpdateSucess()[len(mock.MockUpdateSucess())-1]

		jsonBody, _ := json.Marshal(updateCases.RequestBody)
		req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", "idInvalid"), bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "erroToken"))

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

	t.Run("ErrorNoAuthorization", func(t *testing.T) {
		updateCases := mock.MockUpdateSucess()[len(mock.MockUpdateSucess())-1]

		jsonBody, _ := json.Marshal(updateCases.RequestBody)
		req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

	t.Run("TestingErrosValidation", func(t *testing.T) {

		testeCases := mock.MockUpdateErroValidation()

		for _, testCase := range testeCases {

			jsonBody, _ := json.Marshal(testCase.RequestBody)

			req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			assert.Contains(t, rec.Body.String(), "error")
			TestCount++
		}
	})

}

func TestUpdateErrorDuplication(t *testing.T) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.POST("/api/users", handlers.CreateUserHandle)

	requestBody, _ := mock.MockUserSucess2()

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	t.Run("ErrorDuplicateUpdate", func(t *testing.T) {
		r = gin.New()
		var newUser models.UserResponse

		err = json.Unmarshal(rec.Body.Bytes(), &newUser)

		assert.NoError(t, err)
		jsonBody, _ = json.Marshal(models.UpdateUser{
			Email:    &newUser.Email,
			Nickname: &newUser.Nickname,
		})

		r.PATCH("/api/users/:id", middlewares.AuthMiddleware(), handlers.UpdatedUserHandle)

		req, err = http.NewRequest("PATCH", fmt.Sprintf("/api/users/%v", profileId), bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", goblaToken))

		rec = httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)

		assert.Contains(t, rec.Body.String(), "error")
		TestCount++
	})

}
