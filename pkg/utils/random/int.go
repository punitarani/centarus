package random

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Int64 generates a random int64 number.
func Int64() (int64, error) {
	// Generate a random *big.Int using the crypto/rand package.
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}

	// Convert the *big.Int to an int64 and return it.
	return n.Int64(), nil
}

// Int generates a random int number within a given range.
//
// param min: The minimum value of the number.
// param max: The maximum value of the number.
func Int(min, max int) (int, error) {
	// Generate a random *big.Int within the given range.
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return min, err
	}

	// Add the minimum value to the random *big.Int and return the result as an int.
	return int(n.Int64()) + min, nil
}
