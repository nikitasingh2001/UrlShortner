package helper

import (
	"math/rand"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateUniqueString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func GenRandomString(n int) string {
	return generateUniqueString()
}
