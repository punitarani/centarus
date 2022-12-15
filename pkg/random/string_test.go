package random

import (
	"strings"
	"testing"
)

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
