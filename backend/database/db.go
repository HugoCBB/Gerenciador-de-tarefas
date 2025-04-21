package database

import (
	"fmt"
	"log"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/config"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDataBase() {

	dbConfig := config.NewDatabaseConfigEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err, "\nErro ao conectar ao banco de dados")
	}

	fmt.Println("Conexão ao banco de dados estabelecida com sucesso")

	DB.AutoMigrate(&models.Tarefa{}, &models.User{})
}
