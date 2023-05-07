package database

type Token struct {
	Email  string `json:"email"`
	Token  string `json:"token"`
	Color string `json:"color"`
}
