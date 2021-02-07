package magicbell

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type createUserTest struct {
	name            string
	httpStatus      int
	modifyRequestFn func(*testing.T, *CreateUserRequest)
	checkErr        func(*testing.T, error)
	checkUser       func(*testing.T, *User)
}

func (test createUserTest) Run(t *testing.T, createUserFn func(request CreateUserRequest) (*User, error)) {
	request := initialCreateUserRequest

	if test.modifyRequestFn != nil {
		test.modifyRequestFn(t, &request)
	}

	user, err := createUserFn(request)

	if test.checkErr != nil {
		test.checkErr(t, err)
	}
	if test.checkUser != nil {
		test.checkUser(t, user)
	}
}

type updateUserTest struct {
	name            string
	httpStatus      int
	modifyRequestFn func(*testing.T, *UpdateUserRequest)
	checkErr        func(*testing.T, error)
	checkUser       func(*testing.T, *User)
}

func (test updateUserTest) Run(t *testing.T, userID string, updateUserFn func(userID string, request UpdateUserRequest) (*User, error)) {
	request := initialUpdateUserRequest

	if test.modifyRequestFn != nil {
		test.modifyRequestFn(t, &request)
	}

	user, err := updateUserFn(userID, request)

	if test.checkErr != nil {
		test.checkErr(t, err)
	}
	if test.checkUser != nil {
		test.checkUser(t, user)
	}
}

var (
	createUserTests = []createUserTest{
		{
			name:       "200",
			httpStatus: http.StatusOK,
			checkErr: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
			checkUser: func(t *testing.T, user *User) {
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
			},
		},
		{
			name:       "400",
			httpStatus: http.StatusBadRequest,
			modifyRequestFn: func(t *testing.T, request *CreateUserRequest) {
				request.Email = ""
			},
			checkErr: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.True(t, IsAPIErrors(err))

				errs := err.(APIErrors)
				assert.Len(t, errs, 1)
				assert.Equal(t, APIError{
					Code:    APIErrorCodeUserEmailNotProvided,
					Message: "missing email",
				}, errs[0])
			},
			checkUser: func(t *testing.T, user *User) {
				assert.Nil(t, user)
			},
		},
		{
			name:       "500",
			httpStatus: http.StatusInternalServerError,
			checkErr: func(t *testing.T, err error) {
				require.Error(t, err)

				assert.True(t, IsInternalServerError(err))
				assert.Equal(t, InternalServerError{
					StatusCode: 500,
					Body:       "Internal server error\n",
				}, err.(InternalServerError))
			},
			checkUser: func(t *testing.T, user *User) {
				require.Nil(t, user)
			},
		},
	}
	updateUserTests = []updateUserTest{
		{
			name:       "200",
			httpStatus: http.StatusOK,
			checkErr: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
			checkUser: func(t *testing.T, user *User) {
				require.NotNil(t, user)
				assert.Equal(t, User{
					ID:               "7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c",
					ExternalID:       "56780",
					Email:            "hana@magicbell.io",
					FirstName:        "Hana",
					LastName:         "Mohan",
					CustomAttributes: nil,
				}, *user)
			},
		},
		{
			name:       "400",
			httpStatus: http.StatusBadRequest,
			modifyRequestFn: func(t *testing.T, request *UpdateUserRequest) {
				request.Email = ""
			},
			checkErr: func(t *testing.T, err error) {
				require.Error(t, err)
				assert.True(t, IsAPIErrors(err))

				errs := err.(APIErrors)
				assert.Len(t, errs, 1)
				assert.Equal(t, APIError{
					Code:    APIErrorCodeUserEmailNotProvided,
					Message: "missing email",
				}, errs[0])
			},
			checkUser: func(t *testing.T, user *User) {
				require.Nil(t, user)
			},
		},
		{
			name:       "500",
			httpStatus: http.StatusInternalServerError,
			checkErr: func(t *testing.T, err error) {
				require.Error(t, err)

				assert.True(t, IsInternalServerError(err))
				assert.Equal(t, InternalServerError{
					StatusCode: 500,
					Body:       "Internal server error\n",
				}, err.(InternalServerError))
			},
			checkUser: func(t *testing.T, user *User) {
				require.Nil(t, user)
			},
		},
	}

	initialCreateUserRequest = CreateUserRequest{
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
	initialUpdateUserRequest = UpdateUserRequest{
		ExternalID:       "56780",
		Email:            "hana@magicbell.io",
		FirstName:        "Hana",
		LastName:         "Mohan",
		CustomAttributes: nil,
	}
)

func TestAPI_CreateUser(t *testing.T) {
	for _, test := range createUserTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/users", http.MethodPost, test.httpStatus, func(config Config) {
				api := New(config)
				test.Run(t, api.CreateUser)
			})
		})
	}
}

func TestAPI_UpdateUser(t *testing.T) {
	for _, test := range updateUserTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/users/7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c", http.MethodPut, test.httpStatus, func(config Config) {
				api := New(config)
				test.Run(t, "7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c", api.UpdateUser)
			})
		})
	}
}

func TestCreateUser(t *testing.T) {
	for _, test := range createUserTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/users", http.MethodPost, test.httpStatus, func(config Config) {
				runGlobalTest(config, func() {
					test.Run(t, CreateUser)
				})
			})
		})
	}
}

func TestUpdateUser(t *testing.T) {
	for _, test := range updateUserTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/users/7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c", http.MethodPut, test.httpStatus, func(config Config) {
				runGlobalTest(config, func() {
					test.Run(t, "7fb3ce9f-a866-4dff-8ce8-2f64f7c5ed4c", UpdateUser)
				})
			})
		})
	}
}
