package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password1234"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	login := `password1234`

	logErr := bcrypt.CompareHashAndPassword(bs, []byte(login))
	if logErr != nil {
		fmt.Println(logErr)
	}
}