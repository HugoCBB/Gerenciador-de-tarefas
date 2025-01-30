package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gorilla/mux"
)

func Tarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tarefas)
}

func AdicionarTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	json.NewDecoder(r.Body).Decode(&t)
	database.DB.Create(&t)
	json.NewEncoder(w).Encode(t)
}

func DeletarTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Delete(&t, id)
	json.NewEncoder(w).Encode(t)
}

// func ModificarTarefa(w http.ResponseWriter, r *http.Request) {
// 	var t models.Tarefa

// }
