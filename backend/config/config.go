package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfigEnv struct {
	Host, User, Password, Name, Port string
}

func LoadEnv() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao iniciar arquivo .env")
		return err
	}

	return nil
}

func NewDatabaseConfigEnv() *DatabaseConfigEnv {
	return &DatabaseConfigEnv{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Name:     os.Getenv("DATABASE_NAME"),
		Port:     os.Getenv("DATABASE_PORT"),
	}

}
