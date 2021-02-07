# magicbell-go: Unofficial Go API Library for MagicBell

**This is still a work in progress, expect breaking changes.**

## What is MagicBell?

See https://magicbell.io/.

## Installation

```bash
go get -u github.com/tizz98/magicbell-go
```

## Usage

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
		CustomAttributes: map[string]string{
			"plan":              "enterprise",
			"pricing_version":   "v10",
			"preferred_pronoun": "She",
		},
	})

	fmt.Printf("%#v\n", user)
}
```
