package generator

import (
	"math/rand"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength = 6
)

// GenerateShortCode creates a random short code of a specified length using the defined character set.
func GenerateShortCode() string {
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

// GenerateShortCodeWithSeed creates a random short code using a local rando generator with a specified seed.
func GenerateShortCodeWithSeed(seed int64) string {
	rng := rand.New(rand.NewSource(seed))
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rng.Intn(len(charset))]
	}
	return string(code)
}
