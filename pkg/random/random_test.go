package random

import (
	"math"
	"strings"
	"testing"
)

// TestInt64 tests the Int64 function.
func TestInt64(t *testing.T) {
	// Generate 2**16 random numbers and check if they are within the range.
	for i := 0; i < 1<<16; i++ {
		n, err := Int64()
		if err != nil {
			t.Error(err)
		}
		if n < math.MinInt64 || n > math.MaxInt64 {
			t.Error("generated number is out of range")
		}
	}
}

// TestInt tests the Int function.
func TestInt(t *testing.T) {
	// Generate 2**8 combinations of random numbers and check if they are within the range.
	for i := 0; i < 1<<8; i++ {
		// Generate the lower and upper bounds.
		for l := 0; l < 1<<4; l++ {
			for w := 0; w < 1<<4; w++ {
				min := 1 << (1 << l)
				max := min + (1 << (1 << (w + 1)))

				// Decide the sign of the numbers.
				for s := 0; s < 2; s++ {
					if s == 1 {
						min, max = -max, -min
					}

					// Generate and test the random number.
					n, err := Int(min, max)
					if err != nil {
						t.Error(err)
					}
					if n < min || n > max {
						t.Error("generated number is out of range")
					}
				}
			}
		}
	}
}

// TestString tests the String function.
func TestString(t *testing.T) {
	charsets := []Charset{
		Alpha,
		Digit,
		Special,
		AlphaLower,
		AlphaUpper,
		AlphaNumeric,
		AlphaNumericSpecial,
	}

	// Generate 2**4 strings for each combination and validate them.
	for i := 0; i < 1<<4; i++ {
		// Generate the length and charset.
		for l := 1; l <= 4; l++ {
			n := 1 << (1 << l)
			for _, charset := range charsets {
				// Generate and test for each charset.
				s, err := String(n, charset)
				if err != nil {
					t.Error(err)
				}
				if len(s) != n {
					t.Error("generated string has wrong length")
				}
				for _, c := range s {
					if !strings.Contains(string(charset), string(c)) {
						t.Error("generated string has wrong characters")
					}
				}
			}
		}
	}
}
