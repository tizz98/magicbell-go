package magicbell

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_CreateUser(t *testing.T) {
	request := CreateUserRequest{
		ExternalID: "56780",
		Email:      "hana@magicbell.io",
		FirstName:  "Hana",
		LastName:   "Mohan",
		CustomAttributes: map[string]string{
			"plan":              "enterprise",
			"pricing_version":   "v10",
			"preferred_pronoun": "She",
		},
	}

	t.Run("200", func(t *testing.T) {
		runServer(t, "/users", http.MethodPost, http.StatusOK, func(config Config) {
			api := New(config)
			user, err := api.CreateUser(request)

			require.NoError(t, err)
			require.NotNil(t, user)

			assert.Equal(t, User{
				ID:         "7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c",
				ExternalID: "56780",
				Email:      "hana@magicbell.io",
				FirstName:  "Hana",
				LastName:   "Mohan",
				CustomAttributes: map[string]string{
					"plan":              "enterprise",
					"pricing_version":   "v10",
					"preferred_pronoun": "She",
				},
			}, *user)
		})
	})

	t.Run("400", func(t *testing.T) {
		request := request
		request.Email = ""

		runServer(t, "/users", http.MethodPost, http.StatusBadRequest, func(config Config) {
			api := New(config)
			user, err := api.CreateUser(request)

			require.Error(t, err)
			require.Nil(t, user)

			assert.True(t, IsAPIErrors(err))

			errs := err.(APIErrors)
			assert.Len(t, errs, 1)
			assert.Equal(t, APIError{
				Code:    APIErrorCodeUserEmailNotProvided,
				Message: "missing email",
			}, errs[0])
		})
	})

	t.Run("500", func(t *testing.T) {
		runServer(t, "/users", http.MethodPost, http.StatusInternalServerError, func(config Config) {
			api := New(config)
			user, err := api.CreateUser(request)

			require.Error(t, err)
			require.Nil(t, user)

			assert.True(t, IsInternalServerError(err))
			assert.Equal(t, InternalServerError{
				StatusCode: 500,
				Body:       "Internal server error\n",
			}, err.(InternalServerError))
		})
	})
}

func TestAPI_UpdateUser(t *testing.T) {
	t.Run("200", func(t *testing.T) {

	})

	t.Run("400", func(t *testing.T) {

	})

	t.Run("500", func(t *testing.T) {

	})
}

func TestCreateUser(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}
