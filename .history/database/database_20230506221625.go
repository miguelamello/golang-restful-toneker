package database

import (
	"fmt"
	"time"
)

type Item struct {
	Email string `json:"email"`
	Item string `json:"token"`
	Stamp string `json:"stamp"`
}

var itemMap = make(map[string]Item)

func Timestamp() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func CreateItem( email string) {
	token := Item {
		Email: email, 
		Item: "123456", 
		Stamp: Timestamp(),
	}
	itemMap[email] = token
	fmt.Println(itemMap)
}

func GetItem(email string) string {
	//fmt.Println(email)
	return email
}


