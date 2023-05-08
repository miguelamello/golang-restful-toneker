package database

import (
	"encoding/json"
	"time"
)

type Token struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Stamp string `json:"stamp"`
}

var tokenMap = make(map[string]Token)

func _timestamp() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func _ToJson() string {
	jsonBytes, err := json.Marshal(tokenMap)
	if err != nil { panic(err) }
	return string(jsonBytes)
}

func _createToken( email string) {
	token := Token {
		Email: email, 
		Token: "123456", 
		Stamp: _timestamp(),
	}
	tokenMap[email] = token
}

