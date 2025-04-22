package models

import "errors"

type User struct {
	Id      uint     `json:"id" gorm:"primary_key"`
	Nome    string   `json:"nome"`
	Email   string   `json:"email" gorm:"unique" `
	Senha   string   `json:"senha" `
	Tarefas []Tarefa `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE;"`
}

func (u *User) ValidateRequiredFields() error {
	if u.Email == "" || u.Senha == "" || u.Nome == "" {
		return errors.New("Campos obrigatórios não podem estar vazios")
	}
	return nil
}
