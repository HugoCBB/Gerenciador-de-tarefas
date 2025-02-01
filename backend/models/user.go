package models

type User struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

var users []User
