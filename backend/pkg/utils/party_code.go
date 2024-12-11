package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GeneratePartyCode generates a random alphanumeric code of the given length.
// It ensures thread safety by using a seeded global source for randomness.
func GeneratePartyCode(length int) string {
	if length <= 0 {
		return ""
	}

	// Use a random source for better concurrency handling
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[randSource.Intn(len(charset))]
	}
	return string(code)
}
