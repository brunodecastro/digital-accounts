package util

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	MaybeError(err, "error when encrypt password")
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
