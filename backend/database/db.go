package database

import (
	"fmt"
	"log"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {
	dsn := "hugoADMIN:hugo00028922@tcp(localhost:3306)/tarefas-db"
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err, "\nErro ao conectar ao banco de dados")
	}

	fmt.Println("Conexão ao banco de dados estabelecida com sucesso")

	DB.AutoMigrate(&models.Tarefa{}, &models.User{})
}
