package magicbell

import (
	"context"
	"net/http"
)

// NotificationRecipient is a possible recipient of a notification.
// Generally Email should be specified, but ExternalID can also be provided if the email is not available.
type NotificationRecipient struct {
	// Email is the email of the recipient to send the notification to.
	Email string `json:"email"`
	// ExternalID is the unique string to identify the user in your database.
	ExternalID string `json:"external_id"`
}

// CreateNotificationRequest is the data required to create a new notification
// for a set of recipients in MagicBell. The only required fields are Title and Recipients.
type CreateNotificationRequest struct {
	// Title is the title of the notification
	Title string `json:"title"`
	// Recipients are the users to send the notification to
	Recipients []NotificationRecipient `json:"recipients"`
	// Content is the content of the notification to send
	Content string `json:"content,omitempty"`
	// CustomAttributes are a set of key-value pairs that you can attach to a notification
	CustomAttributes map[string]string `json:"custom_attributes,omitempty"`
	// ActionURL is a URL to redirect the user to when they click on the notification in MagicBell's embeddable notification center.
	ActionURL string `json:"action_url,omitempty"`
	// Category is the category this notification belongs to.
	Category string `json:"category,omitempty"`
}

type createNotificationRequest struct {
	Notification CreateNotificationRequest `json:"notification"`
}

type createNotificationResponse struct {
	baseResponse
	Notification *BaseNotification `json:"notification"`
}

// BaseNotification is the simplest data that can be returned for a MagicBell notification.
// The Notification struct is generally used, but when creating a new notification only the ID is returned.
type BaseNotification struct {
	// ID is the MagicBell ID for the notification
	ID string `json:"id"`
}

// Notification represents a full notification when retrieved from MagicBell.
type Notification struct {
	BaseNotification
	// TODO: more fields
}

// CreateNotification sends a notification to one or multiple users.
func (a *API) CreateNotification(req CreateNotificationRequest) (*BaseNotification, error) {
	return a.CreateNotificationC(context.TODO(), req)
}

// CreateNotification sends a notification to one or multiple users.
func CreateNotification(req CreateNotificationRequest) (*BaseNotification, error) {
	return api.CreateNotification(req)
}

// CreateNotificationC sends a notification to one or multiple users, using a context.Context in the HTTP request.
func (a *API) CreateNotificationC(ctx context.Context, req CreateNotificationRequest) (*BaseNotification, error) {
	var out createNotificationResponse

	if err := a.makeRequest(ctx, http.MethodPost, "notifications", createNotificationRequest{req}, &out); err != nil {
		return nil, err
	}

	return out.Notification, out.Err()
}

// CreateNotificationC sends a notification to one or multiple users, using a context.Context in the HTTP request.
func CreateNotificationC(ctx context.Context, req CreateNotificationRequest) (*BaseNotification, error) {
	return api.CreateNotificationC(ctx, req)
}
