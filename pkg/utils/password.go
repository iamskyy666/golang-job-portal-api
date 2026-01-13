package utils

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateRandomPassword generates a random password of given length
func GenerateRandomPassword(charCount int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	var password strings.Builder

	rand.Seed(time.Now().UnixNano())
	password.Grow(charCount)

	for range charCount {
		password.WriteByte(charset[rand.Intn(len(charset))])
	}

	return password.String()
}
