package database

import (
	"fmt"
	"time"
)

type Ite struct {
	Email string `json:"email"`
	Ite string `json:"token"`
	Stamp string `json:"stamp"`
}

var tokenMap = make(map[string]Ite)

func Timestamp() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func CreateIte( email string) {
	token := Ite {
		Email: email, 
		Ite: "123456", 
		Stamp: Timestamp(),
	}
	tokenMap[email] = token
	fmt.Println(tokenMap)
}

func GetIte(email string) string {
	fmt.Println(email)
	return email
}


