package models

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	DataCriacao string `json:"dataCriacao"`
	// UsuarioID   uint   `json:"usuarioId"`
}

var Tarefas []Tarefa

func (t *Tarefa) ValidateRequiredFields(c *gin.Context) error {
	if t.Titulo == "" || t.Descricao == "" {
		return errors.New("Campos obrigatórios não podem estar vazios")
	}
	return nil

}
