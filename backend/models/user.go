package models

import "errors"

type User struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Nome  string `json:"nome"`
	Email string `json:"email" `
	Senha string `json:"senha" gorm:"size:8"`
	// Tarefas []Tarefa `gorm:"foreignKey:UsuarioID"`
}

var users []User

func (u *User) ValidateRequiredFields() error {
	if u.Email == "" || u.Senha == "" || u.Nome == "" {
		return errors.New("Campos obrigatórios não podem estar vazios")
	}
	return nil
}
