package random

import (
	"crypto/rand"
	"math/big"
)

// Charset is a set of characters that can be used to generate a random string.
type Charset string

const (
	// The Alpha Charset represents the set of uppercase and lowercase letters of the English alphabet.
	Alpha Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// The Digit Charset represents the set of decimal digits.
	Digit Charset = "0123456789"

	// The Special Charset represents a set of special characters commonly used in passwords.
	Special Charset = `!@#$%^&*()_+-=[]{}|;':\",./<>?`

	// The AlphaLower Charset represents the set of lowercase letters of the English alphabet.
	AlphaLower Charset = "abcdefghijklmnopqrstuvwxyz"

	// The AlphaUpper Charset represents the set of uppercase letters of the English alphabet.
	AlphaUpper Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// The AlphaNumeric Charset represents the set of uppercase and lowercase letters and decimal digits
	AlphaNumeric Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// The AlphaNumericSpecial Charset represents the set of uppercase and lowercase, decimal digits, and special characters.
	AlphaNumericSpecial Charset = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;':\",./<>?`
)

// String generate a random string with a given length and charset.
//
// param length: The length of the string.
// param chars: The characters that can be used to generate the string.
func String(length int, chars Charset) (string, error) {
	// Create a slice of bytes to store the generated string.
	var b []byte

	// Generate the random string.
	for i := 0; i < length; i++ {
		// Generate a random integer within the range of the chars slice.
		nb, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))

		if err != nil {
			return "", err
		} else {
			// Append the random character to the slice of bytes.
			b = append(b, chars[nb.Uint64()])
		}
	}

	// Return the generated string.
	return string(b), nil
}
