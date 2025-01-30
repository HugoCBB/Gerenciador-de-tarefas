package models

import "gorm.io/gorm"

type Tarefa struct {
	gorm.Model
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	DataCriacao string `json:"dataCriacao"`
}

var Tarefas []Tarefa
