package magicbell

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type createNotificationTest struct {
	name              string
	httpStatus        int
	modifyRequestFn   func(*testing.T, *CreateNotificationRequest)
	checkErr          func(*testing.T, error)
	checkNotification func(*testing.T, *BaseNotification)
}

func (test createNotificationTest) Run(t *testing.T, createNotificationFn func(CreateNotificationRequest) (*BaseNotification, error)) {
	request := initialCreateNotificationRequest

	if test.modifyRequestFn != nil {
		test.modifyRequestFn(t, &request)
	}

	notification, err := createNotificationFn(request)

	if test.checkErr != nil {
		test.checkErr(t, err)
	}
	if test.checkNotification != nil {
		test.checkNotification(t, notification)
	}
}

var (
	createNotificationTests = []createNotificationTest{
		{
			name:       "201",
			httpStatus: http.StatusCreated,
			checkErr:   assertNoError,
			checkNotification: func(t *testing.T, notification *BaseNotification) {
				require.NotNil(t, notification)
				assert.Equal(t, "ffffff66-ea4f-4da2-afc6-84148b51657a", notification.ID)
			},
		},
		{
			name:       "403",
			httpStatus: http.StatusForbidden,
			checkErr:   assertAPIError(APIErrorCodeForbidden, "not allowed"),
			checkNotification: func(t *testing.T, notification *BaseNotification) {
				require.Nil(t, notification)
			},
		},
		{
			name:       "422",
			httpStatus: http.StatusUnprocessableEntity,
			modifyRequestFn: func(t *testing.T, request *CreateNotificationRequest) {
				request.Recipients = nil
			},
			checkErr: assertAPIError("", `Param 'notification.recipients' is missing`),
			checkNotification: func(t *testing.T, notification *BaseNotification) {
				require.Nil(t, notification)
			},
		},
		{
			name:       "500",
			httpStatus: http.StatusInternalServerError,
			checkErr:   assertInternalServerError,
			checkNotification: func(t *testing.T, notification *BaseNotification) {
				assert.Nil(t, notification)
			},
		},
	}

	initialCreateNotificationRequest = CreateNotificationRequest{
		Title: "Ticket assigned to you: Do you offer demos?",
		Recipients: []NotificationRecipient{{
			Email:      "john@example.com",
			ExternalID: "1924",
		}},
		Content:  "Can I see a demo of your product?",
		Category: "new_message",
		CustomAttributes: map[string]interface{}{
			"order": map[string]string{
				"id":    "12345",
				"title": "A title you can use in your templates",
			},
		},
	}
)

func TestAPI_CreateNotification(t *testing.T) {
	for _, test := range createNotificationTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/notifications", http.MethodPost, test.httpStatus, func(config Config) {
				api := New(config)
				test.Run(t, api.CreateNotification)
			})
		})
	}
}

func TestCreateNotification(t *testing.T) {
	for _, test := range createNotificationTests {
		t.Run(test.name, func(t *testing.T) {
			runServer(t, "/notifications", http.MethodPost, test.httpStatus, func(config Config) {
				runGlobalTest(config, func() {
					Init(config)
					test.Run(t, CreateNotification)
				})
			})
		})
	}
}
