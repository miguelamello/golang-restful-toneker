package database

import (
	"fmt"
	"time"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

var tokenMap = make(map[string]Token)

func timestamp() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func CreateToken( email string) {
	token := Token {
		Email: email, 
		Token: "123456", 
		Stamp: timestamp(),
	}
	tokenMap[email] = token
	fmt.Println(tokenMap)
}


