package random

import (
	"crypto/rand"
	"math"
	"math/big"
)

// RandStringType is a type of random string.
type RandStringType string

// RandStringType Enums
const (
	Alpha               RandStringType = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Digit               RandStringType = "0123456789"
	Special             RandStringType = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
	AlphaLower          RandStringType = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpper          RandStringType = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNumeric        RandStringType = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	AlphaNumericSpecial RandStringType = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;':\",./<>?"
)

// Int64 returns a random int64 number.
func Int64() (int64, error) {
	// Generate a random *big.Int using the crypto/rand package.
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}

	// Convert the *big.Int to an int64 and return it.
	return n.Int64(), nil
}

// Int returns a random int number.
//
// min is the minimum number.
// max is the maximum number.
func Int(min, max int) (int, error) {
	// Generate a random *big.Int within the given range.
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return min, err
	}

	// Add the minimum value to the random *big.Int and return the result as an int.
	return int(n.Int64()) + min, nil
}

// String returns a random string.
//
// length is the length of the string.
// randStringType is the type of random string.
func String(length int, chars RandStringType) (string, error) {
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
