package utils

import (
	"crypto/sha256"
	"math/big"
)

// SHA-256 hashing to generate a unique number for the URL
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base62Encode converts a number into a Base62 string (a-z, A-Z, 0-9)
func base62Encode(number uint64) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := uint64(len(alphabet))
	var encoded string

	for number > 0 {
		encoded = string(alphabet[number%length]) + encoded
		number = number / length
	}
	return encoded
}

// GenerateShortLink is the main function we will call
func GenerateShortLink(initialUrl string) string {
	urlHashBytes := sha256Of(initialUrl)
	// Convert the hash bytes into a large number
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base62Encode(generatedNumber)
	return finalString[0:8] // Return only the first 8 characters
}
