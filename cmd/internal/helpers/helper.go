package helpers

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomNumbers(length int) (string, error) {

	const numbers = "1234567890"

	buffer := make([]byte, length)

	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	numLength := len(numbers)

	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%numLength]
	}

	return string(buffer), nil
}

func GenerateCode() (string, error) {
	return RandomNumbers(6)
}
