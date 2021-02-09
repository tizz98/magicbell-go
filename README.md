# magicbell-go: Unofficial Go API Library for MagicBell

[![Go Reference](https://pkg.go.dev/badge/github.com/tizz98/magicbell-go.svg)](https://pkg.go.dev/github.com/tizz98/magicbell-go)
[`mbctl`](#mbctl-cli-installation)

**This is still a work in progress, expect breaking changes until v1.0.0**

## What is MagicBell?

See https://magicbell.io/.

## Library Installation

```bash
go get -u github.com/tizz98/magicbell-go
```

## Library Usage

### Send Notification

```go
package main

import (
	"fmt"

	"github.com/tizz98/magicbell-go"
)

func main() {
	magicbell.Init(magicbell.Config{
		APIKey:    "my-key",
		APISecret: "my-secret",
	})

	notification, _ := magicbell.CreateNotification(magicbell.CreateNotificationRequest{
		Title: "Welcome to MagicBell",
		Recipients: []magicbell.NotificationRecipient{
			{Email: "hana@magicbell.io"},
			{ExternalID: "some-id"},
        },
		Content: "The notification inbox for your product. Get started in minutes.",
		CustomAttributes: map[string]interface{}{
			"order": map[string]string{
				"id": "1234567",
				"title": "A title you can use in your templates",
			},
		},
		ActionURL: "https://developer.magicbell.io",
		Category:  "new_message",
	})
	fmt.Printf("Created notification %s\n", notification.ID)
}
```

### Create user

```go
package main

import (
	"fmt"
	
	"github.com/tizz98/magicbell-go"
)

func main() {
	magicbell.Init(magicbell.Config{
		APIKey:    "my-key",
		APISecret: "my-secret",
	})

	user, _ := magicbell.CreateUser(magicbell.CreateUserRequest{
		ExternalID: "56780",
		Email:      "hana@magicbell.io",
		FirstName:  "Hana",
		LastName:   "Mohan",
		CustomAttributes: map[string]interface{}{
			"plan":              "enterprise",
			"pricing_version":   "v10",
			"preferred_pronoun": "She",
		},
	})

	fmt.Printf("%#v\n", user)
}
```


## `mbctl` CLI Installation

Download the latest release for your OS from https://github.com/tizz98/magicbell-go/releases 
and add the binary to your `PATH`.

## `mbctl` CLI Usage

```bash
mbctl --version
mbctl --help
```

### Initialize Config

This will save your API key and API secret in a `config.yaml` file.

```bash
mbctl config init
```

### Notification Commands

Commands related to Notifications.

#### Create Notification

Send a notification to a set of users. Separate multiple recipients with a comma.
Any string with an `@` will be considered an email address and sent to the API in that field.

```bash
mbctl notifications create \
  --title="CLI test" \
  --recipients hana@magicbell.io,foo@example.com,my-external-id \
  --content="Notification content" \
  --action-url https://google.com \
  --category new_message
```

### User Commands

Commands related to Users.

#### Generate HMAC

Generate and return a base64-encoded HMAC signature of the provided email.
The HMAC key is the API secret.

```bash
mbctl users generate-hmac hana@magicbell.io
```
