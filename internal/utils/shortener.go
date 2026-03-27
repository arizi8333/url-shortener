package utils

import (
	"math/rand"
	"time"
)

// charset contains the characters used for generating short codes.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateShortCode generates a random short code of the specified length.
func GenerateShortCode(length int) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time.

	// Generate a random short code by selecting characters from the charset.
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
