package main

import (
	"backend-challenge/internal/validator"
	"fmt"
)

func main() {
	password := "Abcdefgh1!"
	isValid := validator.PasswordIsValid(password)
	fmt.Printf("Password %s is valid: %v\n", password, isValid)
}
