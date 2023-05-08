/*
	Package database implements a simple in-memory database	
	like Redis or Memcached. It is used to store the token 
	for fast input and output. It is not persistent, so 
	in the background it is necessary to implement a routine 
	to save data to a persistent database like MongoDB.
*/

package database

import (
	//"fmt"
	"encoding/json"
	"time"
	"crypto/rand"
    "encoding/base64"
)

// Strut simulating an item in the database
type Item struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

// Map simulating the database collection
var itemMap = make(map[string]Item)

// Return the token related to the user mail
func GetPayload(user string) string {
	item := itemMap[user]
	l := len(item.Email)
	if l > 2 {
		jsonBytes, _ := json.Marshal(item)
		jsonString := string(jsonBytes)
		return jsonString
	} else {
		return "{}"
	}
}

// Get the current date and time
func GetCreated() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

// Get the expiration date of the token
func GetExpires(timestamp string) string {
	return timestamp
}

// Create a new item in the database
func CreateItem(email string) {
	stamp := GetCreated()
	token := Item{
		Email:   email,
		Token:   "123456",
		Created: stamp,
		Expires: GetExpires(stamp),
	}
	itemMap[email] = token
	//fmt.Println(itemMap)
}

// Get an item from the database
func GetItem(email string) string {
	//fmt.Println(email)
	return email
}
