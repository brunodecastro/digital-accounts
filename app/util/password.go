package util

import "golang.org/x/crypto/bcrypt"

// EncryptPassword - returns the bcrypt hash of the password
func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	MaybeError(err, "error when encrypt password")
	return string(hash)
}

// CheckPasswordHash - compares a bcrypt hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
