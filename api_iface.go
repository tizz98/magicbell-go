package magicbell

import "context"

// IAPI contains all the methods for working with the MagicBell API
type IAPI interface {
	// GenerateUserEmailHMAC generates a sha256 HMAC signature of the user's email
	// using the APISecret as the HMAC key. The returned value is a base64 encoded
	// string of the resulting HMAC signature. See https://developer.magicbell.io/reference#performing-api-requests-from-javascript
	GenerateUserEmailHMAC(userEmail string) string
	//CreateNotification()
	//CreateNotificationC(ctx context.Context)
	//FetchUserNotifications()
	//FetchUserNotificationsC(ctx context.Context)
	//FetchUserNotification()
	//FetchUserNotificationC(ctx context.Context)
	//DeleteUserNotification()
	//DeleteUserNotificationC(ctx context.Context)
	//MarkNotificationRead()
	//MarkNotificationReadC(ctx context.Context)
	//MarkNotificationUnread()
	//MarkNotificationUnreadC(ctx context.Context)
	//MarkAllNotificationsRead()
	//MarkAllNotificationsReadC(ctx context.Context)
	//MarkAllNotificationsSeen()
	//MarkAllNotificationsSeenC(ctx context.Context)
	// CreateUser creates a new user in MagicBell.
	// Please note that you must provide the user's email or the external id so MagicBell can uniquely identify the user.
	// The external id, if provided, must be unique to the user.
	CreateUser(req CreateUserRequest) (*User, error)
	// CreateUserC creates a new user in MagicBell, using a context.Context in the HTTP request.
	// Please note that you must provide the user's email or the external id so MagicBell can uniquely identify the user.
	// The external id, if provided, must be unique to the user.
	CreateUserC(ctx context.Context, req CreateUserRequest) (*User, error)
	// UpdateUser updates a user in MagicBell with the given ID.
	// The user id is the MagicBell user id. Alternatively, provide an id like
	// email:theusersemail@example.com or external_id:theusersexternalid as the user id.
	UpdateUser(userID string, req UpdateUserRequest) (*User, error)
	// UpdateUserC updates a user in MagicBell with the given ID, using a context.Context in the HTTP request.
	// The user id is the MagicBell user id. Alternatively, provide an id like
	// email:theusersemail@example.com or external_id:theusersexternalid as the user id.
	UpdateUserC(ctx context.Context, userID string, req UpdateUserRequest) (*User, error)
}
