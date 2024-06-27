package auth

import "golang.org/x/crypto/bcrypt"


func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}

	return string(hash), nil
}