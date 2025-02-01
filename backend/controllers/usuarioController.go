package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
)

func Usuarios(w http.ResponseWriter, r *http.Request) {
	var u []models.User
	database.DB.Find(&u)
	json.NewEncoder(w).Encode(u)

}

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	database.DB.Create(&u)
	json.NewEncoder(w).Encode(u)

}
