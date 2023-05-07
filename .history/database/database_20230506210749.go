package database

import (
	"encoding/json"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

var tokenMap = make(map[string]Token)

func createToken( email string) string {
	token := Token {
		Email: email, 
		Token: ""
	}
}


