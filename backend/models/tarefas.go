package models

type Tarefa struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	DataCriacao string `json:""`
}

var Tarefas []Tarefa
