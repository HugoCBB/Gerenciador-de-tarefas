package main

import (
	"fmt"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/config"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/routers"
)

func init() {
	config.LoadEnv()
	database.ConectDataBase()
}

func main() {
	routers.HandleRequest()
	fmt.Println("Servidor rodando na porta 8000")

}
