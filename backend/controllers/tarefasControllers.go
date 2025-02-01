package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gorilla/mux"
)

func Tarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database.DB.Find(&models.Tarefas)
	json.NewEncoder(w).Encode(models.Tarefas)
}

func AdicionarTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	json.NewDecoder(r.Body).Decode(&t)
	t.DataCriacao = time.Now().String()
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

func ModificarTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&t, id)
	json.NewDecoder(r.Body).Decode(&t)
	database.DB.Save(&t)
	json.NewEncoder(w).Encode(t)
}

func ObterTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&t, id)
	json.NewEncoder(w).Encode(t)
}
