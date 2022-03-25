package core

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(text string, publicKey int) string {

	b, _ := bcrypt.GenerateFromPassword([]byte(text), publicKey)
	return string(b)
}
