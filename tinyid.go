package tinyid

import (
	"crypto/rand"
	"encoding/base64"
)

// New generates a new tiny id based on specified length
func New(length int) string {
	idLen := length
	if idLen == 0 {
		idLen = 2
	}

	idLen += 4

	b := make([]byte, idLen)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	base64Encoded := base64.RawURLEncoding.EncodeToString(b)

	// Truncate the string to the desired length
	if len(base64Encoded) > length && length > 0 {
		base64Encoded = base64Encoded[:length]
	}

	return base64Encoded
}
