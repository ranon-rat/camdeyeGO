package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123456"

	b, _ := bcrypt.GenerateFromPassword([]byte(password), rand.Intn(10))
	c, err := bcrypt.GenerateFromPassword([]byte(password), rand.Intn(10))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string((b)), string(c))
}
