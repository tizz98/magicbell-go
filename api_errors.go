package magicbell

import (
	"fmt"
	"net/http"
)

// APIErrorCode is a string returned from the API that represents
// some known error state. See https://developer.magicbell.io/reference#error-codes
type APIErrorCode string

const (
	// APIErrorCodeAPIKeyNotProvided is returned when no API key was provided with the HTTP request
	APIErrorCodeAPIKeyNotProvided APIErrorCode = "api_key_not_provided"
	// APIErrorCodeIncorrectAPIKey is returned when an API key was provided but is not valid
	APIErrorCodeIncorrectAPIKey APIErrorCode = "incorrect_api_key"
	// APIErrorCodeAPISecretNotProvided is returned when no API secret was provided with the HTTP request
	APIErrorCodeAPISecretNotProvided APIErrorCode = "api_secret_not_provided"
	// APIErrorCodeAPISecretIsIncorrect is returned when an API secret was provided but is not valid
	APIErrorCodeAPISecretIsIncorrect APIErrorCode = "api_secret_is_incorrect"
	// APIErrorCodeForbidden is returned when the API request is forbidden because of a permission issue
	APIErrorCodeForbidden APIErrorCode = "forbidden"
	// APIErrorCodeNeitherUserHMACNorAPISecretProvided is returned when neither the User email HMAC
	// nor the API secret are not provided. See https://developer.magicbell.io/reference#performing-api-requests-from-javascript
	APIErrorCodeNeitherUserHMACNorAPISecretProvided APIErrorCode = "neither_user_hmac_nor_api_secret_provided"
	// APIErrorCodeUserEmailNotProvided is returned when a user's email should have been provided with the HMAC
	// but was not set.
	APIErrorCodeUserEmailNotProvided APIErrorCode = "user_email_not_provided"
)

// APIError represents a single error returned from the
// MagicBell API. It's unclear which fields are mandatory to be returned,
// so only count on APIError.Code being set and check the other fields
// for being empty strings before use.
type APIError struct {
	Code       APIErrorCode `json:"code"`
	Suggestion string       `json:"suggestion"`
	Message    string       `json:"message"`
	HelpLink   string       `json:"help_link"`
}

// Error returns the message contained in this APIError
// or the APIErrorCode if no message is set.
func (e APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return string(e.Code)
}

// IsCode returns true when the provided APIErrorCode
// matches the stored APIError.Code
func (e APIError) IsCode(c APIErrorCode) bool {
	return e.Code == c
}

// APIErrors is an alias to a slice of APIError.
// This is useful for parsing API responses.
type APIErrors []APIError

// Error returns a string of the first APIError
func (e APIErrors) Error() string {
	if len(e) > 0 {
		return e[0].Error()
	}
	return ""
}

// IsAPIErrors returns true when err is not nil and
// is an instance of APIErrors
func IsAPIErrors(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(APIErrors)
	return ok
}

// InternalServerError represents a 5xx HTTP error returned from the API
type InternalServerError struct {
	// StatusCode is the HTTP status code returned from the request
	StatusCode int
	// Body is the response body, if any
	Body string
}

// Error returns a simple string describing the HTTP error
func (e InternalServerError) Error() string {
	return fmt.Sprintf("HTTP %d %s", e.StatusCode, http.StatusText(e.StatusCode))
}

// IsInternalServerError returns true when the underlying error is an InternalServerError
func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(InternalServerError)
	return ok
}
