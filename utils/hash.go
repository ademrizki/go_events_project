package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPasswordInBytes), err

}

func VerifyPassword(hashedPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
