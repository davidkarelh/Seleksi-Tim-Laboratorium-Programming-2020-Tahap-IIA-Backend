package helper

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash a pasword")
	}
	return string(hash)
}
