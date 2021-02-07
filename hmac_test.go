package magicbell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var hmacTests = []struct {
	name         string
	config       Config
	userEmail    string
	expectedHMAC string
}{
	{
		name:         "happy case with valid data",
		config:       Config{APISecret: "secret"},
		userEmail:    "mary@example.com",
		expectedHMAC: "0FoKrRrv40mSiO+WjHaTw/F/71fxEY57pS98r5uK4DE=",
	},
	{
		name:         "missing api secret",
		config:       Config{},
		userEmail:    "mary@example.com",
		expectedHMAC: "27onq9MhtVBlE1HgplRVPMNIzZUuy1O5hFuT2j18bXI=",
	},
	{
		name:         "empty email",
		config:       Config{APISecret: "secret"},
		userEmail:    "",
		expectedHMAC: "+eZuF5tnR65UEI+C+K3os8Jddv0wr95sOVgixTAZYWk=",
	},
}

func TestAPI_GenerateUserEmailHMAC(t *testing.T) {
	for _, test := range hmacTests {
		t.Run(test.name, func(t *testing.T) {
			api := New(test.config)
			assert.Equal(t, test.expectedHMAC, api.GenerateUserEmailHMAC(test.userEmail))
		})
	}
}

func TestGenerateUserEmailHMAC(t *testing.T) {
	for _, test := range hmacTests {
		runGlobalTest(test.config, func() {
			t.Run(test.name, func(t *testing.T) {
				assert.Equal(t, test.expectedHMAC, GenerateUserEmailHMAC(test.userEmail))
			})
		})
	}
}
