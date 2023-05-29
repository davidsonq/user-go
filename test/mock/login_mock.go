package mock

import (
	"github.com/davidsonq/user-go/internal/models"
)

type TestCasesLogin struct {
	RequestBody models.LoginUser
}

func MockUserSucessLogin() (models.LoginUser, models.LoginResponse) {
	return models.LoginUser{
			Email:    "test@example.com",
			Password: "123456789",
		}, models.LoginResponse{
			Token: "text",
		}
}

func MockUserForbiden() []TestCasesLogin {
	testeCases := []TestCasesLogin{
		{
			// Error Email invalid
			RequestBody: models.LoginUser{
				Email:    "testerr@example.com",
				Password: "123456789",
			},
		},
		{
			// Error Password invalid
			RequestBody: models.LoginUser{
				Email:    "test@example.com",
				Password: "12345676",
			},
		},
	}
	return testeCases
}

func MockUsersErrorValidationLogin() []TestCasesLogin {

	testeCases := []TestCasesLogin{

		{
			// Error Email required
			RequestBody: models.LoginUser{
				Password: "123456789",
			},
		},
		{
			// Error Password required
			RequestBody: models.LoginUser{
				Email: "test@example.com",
			},
		},
		{
			// Error Password min
			RequestBody: models.LoginUser{
				Email:    "test@example.com",
				Password: "1",
			},
		},
		{
			// Error Password max
			RequestBody: models.LoginUser{
				Email:    "test@example.com",
				Password: "122245645645645445645646545646546545656456456456446546",
			},
		},
		{
			// Error Email max
			RequestBody: models.LoginUser{
				Email:    "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest@example.com",
				Password: "123456",
			},
		},
		{
			// Error Email invalid
			RequestBody: models.LoginUser{
				Password: "123456",
				Email:    "test",
			},
		},
	}

	return testeCases
}
