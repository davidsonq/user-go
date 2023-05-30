package mock

import (
	"github.com/davidsonq/user-go/internal/models"
	"github.com/google/uuid"
)

type TestCases struct {
	RequestBody models.User
}
type UpdateCases struct {
	RequestBody models.UpdateUser
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

func MockUserSucess2() (models.User, models.User) {
	requestBody := models.User{
		ID:       uuid.New().String(),
		Nickname: "test1",
		Email:    "test1@example.com",
		Password: "123456789",
	}

	response := models.User{
		Nickname: "test1",
		Email:    "test1@example.com",
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

func MockUpdateSucess() []UpdateCases {
	email := "test1@example.com"
	email2 := "test@example.com"

	password := "550123"
	password2 := "123456789"

	nickname := "testupdate"
	nickname2 := "test"

	requestBody := []UpdateCases{
		{
			RequestBody: models.UpdateUser{
				Email:    &email,
				Nickname: &nickname2,
				Password: &password2,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Email:    &email2,
				Nickname: &nickname2,
				Password: &password,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Email:    &email2,
				Password: &password2,
				Nickname: &nickname,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Email:    &email2,
				Password: &password2,
				Nickname: &nickname2,
			},
		},
	}

	return requestBody
}

func MockUpdateErroValidation() []UpdateCases {
	passwordMin := "1"
	passwordMax := "12156d4fa54d4a6sdf654asd4f56a4sdfas4fda4f4da4dsf564adsf4564adsf564asfd564ds"

	emailInvalid := "test"
	emailMax := "testadsfadfasdfasdfasdfasdfasdfasdfasfasdfasdfasdfasdfasfdasdfasdfafdsasdfa@hotmail.com"

	nicknameMin := "t"
	nicknameMax := "eidfkadfashuasoudfoasdhfoiahsdoifhaosidhfoiashdfoiashdfoiahsdoifhasoidfhaosidhfoiashdfoiahsdhsfadoi"

	requestBody := []UpdateCases{
		{
			RequestBody: models.UpdateUser{
				Email: &emailInvalid,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Email: &emailMax,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Password: &passwordMin,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Password: &passwordMax,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Nickname: &nicknameMin,
			},
		},
		{
			RequestBody: models.UpdateUser{
				Nickname: &nicknameMax,
			},
		},
	}
	return requestBody
}
