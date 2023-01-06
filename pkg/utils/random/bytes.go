package random

import (
	"crypto/rand"
)

func Bytes(n uint) ([]byte, error) {
	// Make a new slice of length n
	b := make([]byte, n)

	// Fill the slice with random bytes
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, err
}
