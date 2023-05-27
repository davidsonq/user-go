package mock

import (
	"user-go/models"

	"github.com/google/uuid"
)

type TestCases struct {
	RequestBody models.User
}

func MockUserSucess() (models.User, models.User) {
	requestBody := models.User{
		ID:       uuid.New().String(),
		Nickname: "test",
		Email:    "test@example.com",
		Password: "123456789",
	}

	response := models.User{
		Nickname: "test",
		Email:    "test@example.com",
	}

	return requestBody, response
}

func MockUserDuplicate() []TestCases {

	testeCases := []TestCases{

		// Duplicate email
		{
			models.User{
				ID:       uuid.New().String(),
				Nickname: "test2",
				Email:    "test@example.com",
				Password: "123456789",
			},
		},

		// Duplicate NickName
		{
			models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "test2@example.com",
				Password: "123456789",
			},
		},
	}
	return testeCases
}

func MockUsersErrorValidation() []TestCases {
	testeCases := []TestCases{
		{
			// Error NickName required
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Email:    "test@example.com",
				Password: "123456789",
			},
		},
		{
			// Error Email required
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Password: "123456789",
			},
		},
		{
			// Error Password required
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "test@example.com",
			},
		},
		{
			// Error Password min
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "test@example.com",
				Password: "1",
			},
		},
		{
			// Error Password max
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "test@example.com",
				Password: "122245645645645445645646545646546545656456456456446546",
			},
		},
		{
			// Error Email max
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest@example.com",
				Password: "123456",
			},
		},
		{
			// Error Email invalid
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "test",
				Email:    "test",
				Password: "123456",
			},
		},
		{
			// Error NickName min
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "te",
				Email:    "test@exemple",
				Password: "123456",
			},
		},
		{
			// Error NickName max
			RequestBody: models.User{
				ID:       uuid.New().String(),
				Nickname: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				Email:    "test@exemple",
				Password: "123456",
			},
		},
	}

	return testeCases
}
