package main

import (
	"fmt"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/config"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/routers"
)

func init() {
	config.LoadEnv()
}

func main() {

	database.ConectDataBase()
	fmt.Println("Servidor rodando na porta 8000")
	routers.HandleRequest()

}
