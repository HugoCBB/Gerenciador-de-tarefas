package main

import (
	"fmt"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/routers"
)

func main() {
	database.ConectDataBase()
	fmt.Println("Servidor rodando na porta 8000")
	// models.Tarefas = []models.Tarefa{
	// 	{ID: 1, Titulo: "Estudar Go", Descricao: "Estudar Go para melhorar o conhecimento", DataCriacao: "2021-07-01"},
	// 	{ID: 2, Titulo: "Estudar React", Descricao: "Estudar React para melhorar o conhecimento", DataCriacao: "2021-07-01"},
	// }
	routers.HandleRequest()

}
