package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error occured while hashing password ->", err)
	}

	return string(hashedPwd)
}

func ComparePassword(pwd, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}
