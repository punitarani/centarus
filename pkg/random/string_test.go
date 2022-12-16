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

	for n := 8; n <= 4092; n <<= 1 {
		for _, charset := range charsets {
			for i := 0; i < 1<<4; i++ {
				// Generate and test for each length and charset 2**4 times.
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
