package database

import (
	"fmt"
	"log"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {

	dsn := "host=localhost user=postgres password=hugo00028922 dbname=tarefas-db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err, "\nErro ao conectar ao banco de dados")
	}

	fmt.Println("Conexão ao banco de dados estabelecida com sucesso")

	DB.AutoMigrate(&models.Tarefa{}, &models.User{})
}
