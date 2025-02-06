package models

import "time"

type Tarefa struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	DataCriacao time.Time `json:"dataCriacao"`
	// UsuarioID   uint   `json:"usuarioId"`
}

var Tarefas []Tarefa
