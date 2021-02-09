package magicbell

import (
	"context"
	"fmt"
	"net/http"
)

// CreateUserRequest is the set of data required to create a new user in MagicBell.
// Please note that you must provide the user's email or the external id so MagicBell can uniquely identify the user.
// The external id, if provided, must be unique to the user.
type CreateUserRequest struct {
	// ExternalID is a unique string that MagicBell can utilize to uniquely identify the user.
	// We recommend setting this attribute to the ID of the user in your database.
	// Provide the external id if the user's email is unavailable.
	ExternalID string `json:"external_id"`
	// Email is the user's email.
	Email string `json:"email"`
	// FirstName is the user's first name.
	FirstName string `json:"first_name,omitempty"`
	// LastName is the user's last name.
	LastName string `json:"last_name,omitempty"`
	// CustomAttributes are any customers attributes that you'd like to associate with the user.
	// These custom attributes can later be utilized in MagicBell's web interface (when writing email templates for example).
	CustomAttributes CustomAttributes `json:"custom_attributes,omitempty"`
}

// UpdateUserRequest is the set of data required to update an existing user in MagicBell.
// Note that this is a PUT request, so all data given will overwrite all existing data in MagicBell.
type UpdateUserRequest struct {
	// ExternalID is a unique string that MagicBell can utilize to uniquely identify the user.
	// We recommend setting this attribute to the ID of the user in your database.
	// Provide the external id if the user's email is unavailable.
	ExternalID string `json:"external_id"`
	// Email is the user's email.
	Email string `json:"email"`
	// FirstName is the user's first name.
	FirstName string `json:"first_name,omitempty"`
	// LastName is the user's last name.
	LastName string `json:"last_name,omitempty"`
	// CustomAttributes are any customers attributes that you'd like to associate with the user.
	// These custom attributes can later be utilized in MagicBell's web interface (when writing email templates for example).
	CustomAttributes CustomAttributes `json:"custom_attributes,omitempty"`
}

// User is MagicBell's representation of a user, uniquely identified by ID.
type User struct {
	// ID is the unique ID MagicBell internally uses to represent your user.
	ID string `json:"id"`
	// ExternalID is a unique string that MagicBell can utilize to uniquely identify the user.
	ExternalID string `json:"external_id"`
	// Email is the user's email.
	Email string `json:"email"`
	// FirstName is the user's first name.
	FirstName string `json:"first_name"`
	// LastName is the user's last name.
	LastName string `json:"last_name"`
	// CustomAttributes are any customers attributes that you'd like to associate with the user.
	CustomAttributes CustomAttributes `json:"custom_attributes"`
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
