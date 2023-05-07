package database

import (
	"encoding/json"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

func createToken( email string) string {
	return ""
}

var tokenMap map[string]Token
tokenMap["mytoken"] = Token

