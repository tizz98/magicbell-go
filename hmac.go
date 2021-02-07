package magicbell

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateUserEmailHMAC generates a sha256 HMAC signature of the user's email
// using the APISecret as the HMAC key. The returned value is a base64 encoded
// string of the resulting HMAC signature. See https://developer.magicbell.io/reference#performing-api-requests-from-javascript
func (a *API) GenerateUserEmailHMAC(userEmail string) string {
	mac := hmac.New(sha256.New, []byte(a.config.APISecret))
	mac.Write([]byte(userEmail))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// GenerateUserEmailHMAC is a global shortcut to API.GenerateUserEmailHMAC
func GenerateUserEmailHMAC(userEmail string) string { return api.GenerateUserEmailHMAC(userEmail) }
