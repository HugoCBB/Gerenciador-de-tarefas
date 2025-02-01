package main

import (
	"fmt"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/routers"
)

func main() {
	database.ConectDataBase()
	fmt.Println("Servidor rodando na porta 8000")
	routers.HandleRequest()

}
