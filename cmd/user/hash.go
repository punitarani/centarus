package user

import "golang.org/x/crypto/argon2"

const KeyLen = 32

// HashPassword hashes a password using Argon2id.
//
// param password: The password to hash.
// param salt: The salt to use.
// return: The hashed password bytes.
func HashPassword(password string, salt []byte) []byte {
	// Convert to bytes
	passwordBytes := []byte(password)

	// Generate the Argon2id hash of the password
	hash := argon2.IDKey(passwordBytes, salt, 16, 64*1024, 4, KeyLen)

	return hash
}
