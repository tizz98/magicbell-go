package magicbell

import (
	"context"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	ExternalID       string            `json:"external_id"`
	Email            string            `json:"email"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	CustomAttributes map[string]string `json:"custom_attributes"`
}

type UpdateUserRequest struct {
	ExternalID       string            `json:"external_id"`
	Email            string            `json:"email"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	CustomAttributes map[string]string `json:"custom_attributes"`
}

type User struct {
	ID               string            `json:"id"`
	ExternalID       string            `json:"external_id"`
	Email            string            `json:"email"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	CustomAttributes map[string]string `json:"custom_attributes"`
}

type createUserRequest struct {
	User CreateUserRequest `json:"user"`
}

type createUserResponse struct {
	baseResponse
	User *User `json:"user"`
}

type updateUserRequest struct {
	User UpdateUserRequest `json:"user"`
}

type updateUserResponse struct {
	baseResponse
	User *User `json:"user"`
}

// CreateUser creates a new user in MagicBell.
// Please note that you must provide the user's email or the external id so MagicBell can uniquely identify the user.
// The external id, if provided, must be unique to the user.
func (a *API) CreateUser(req CreateUserRequest) (*User, error) {
	return a.CreateUserC(context.TODO(), req)
}

// CreateUser is a global shortcut to API.CreateUser
func CreateUser(req CreateUserRequest) (*User, error) { return api.CreateUser(req) }

// CreateUserC creates a new user in MagicBell, using a context.Context in the HTTP request.
// Please note that you must provide the user's email or the external id so MagicBell can uniquely identify the user.
// The external id, if provided, must be unique to the user.
func (a *API) CreateUserC(ctx context.Context, req CreateUserRequest) (*User, error) {
	var out createUserResponse

	if err := a.makeRequest(ctx, http.MethodPost, "users", createUserRequest{req}, &out); err != nil {
		return nil, err
	}

	return out.User, out.Err()
}

// CreateUserC is a global shortcut to API.CreateUserC
func CreateUserC(ctx context.Context, req CreateUserRequest) (*User, error) {
	return api.CreateUserC(ctx, req)
}

// UpdateUser updates a user in MagicBell with the given ID.
// The user id is the MagicBell user id. Alternatively, provide an id like
// email:theusersemail@example.com or external_id:theusersexternalid as the user id.
func (a *API) UpdateUser(userID string, req UpdateUserRequest) (*User, error) {
	return a.UpdateUserC(context.TODO(), userID, req)
}

// UpdateUser is a global shortcut to API.UpdateUser
func UpdateUser(userID string, req UpdateUserRequest) (*User, error) {
	return api.UpdateUser(userID, req)
}

// UpdateUserC updates a user in MagicBell with the given ID, using a context.Context in the HTTP request.
// The user id is the MagicBell user id. Alternatively, provide an id like
// email:theusersemail@example.com or external_id:theusersexternalid as the user id.
func (a *API) UpdateUserC(ctx context.Context, userID string, req UpdateUserRequest) (*User, error) {
	var out updateUserResponse

	if err := a.makeRequest(ctx, http.MethodPut, fmt.Sprintf("users/%s", userID), updateUserRequest{req}, &out); err != nil {
		return nil, err
	}

	return out.User, out.Err()
}

// UpdateUserC is a global shortcut to API.UpdateUserC
func UpdateUserC(ctx context.Context, userID string, req UpdateUserRequest) (*User, error) {
	return api.UpdateUserC(ctx, userID, req)
}
