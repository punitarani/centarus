package random

import "testing"

func TestBytes(t *testing.T) {
	// Run each length 2**4 times
	for i := 0; i < 1<<4; i++ {
		// Generate the length.
		for l := 1; l <= 1<<15; l++ {
			n, err := Bytes(uint(l))
			if err != nil {
				t.Error(err)
			}
			if len(n) != l {
				t.Error("generated number is out of range")
			}
		}
	}
}
