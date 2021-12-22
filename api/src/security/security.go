package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash into it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare a password and a hash and return if they are equal
func VerifyPassword(passwordWithHash, passwordWithoutHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordWithoutHash))
}
