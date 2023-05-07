package database

import (
	"fmt"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

func createToken( email string) string {
	return ""
}

tokenMap := make(map[string]Token)
tokenMap