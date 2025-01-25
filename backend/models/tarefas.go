package models

type Tarefa struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	DataCriacao string `json:"dataCriacao"`
}

var Tarefas []Tarefa
