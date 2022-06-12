package models

type User struct {
	Name   string `json:"name"`
	Genger string `json:"gender"`
	Age    int    `json:"age"`
}

type UserEntity struct {
	Entity
	User
}
