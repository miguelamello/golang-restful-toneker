package database

import (
	"fmt"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

func (t *Token) create() string {
	return t.Token
}