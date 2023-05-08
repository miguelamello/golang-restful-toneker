package database

import (
	"fmt"
	"time"
)

type Item struct {
	Email string `json:"email"`
	Item string `json:"token"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

var itemMap = make(map[string]Item)

func GetCreated() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}



func CreateItem( email string) {
	token := Item {
		Email: email, 
		Item: "123456", 
		Created: GetCreated(),
		Expires: Timestamp(),
	}
	itemMap[email] = token
	fmt.Println(itemMap)
}

func GetItem(email string) string {
	//fmt.Println(email)
	return email
}


