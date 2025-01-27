package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
)

func Tarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Tarefas)
}

func AdicionarTarefa(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	json.NewDecoder(r.Body).Decode(&t)
	models.Tarefas = append(models.Tarefas, t)
	json.NewEncoder(w).Encode(t)
}

// func DeletarTarefa(w http.ResponseWriter, r *http.Request) {
// 	var t models.Tarefa
// 	json.NewDecoder(r.Body).Decode(&t)

// 	for i, tarefa := range models.Tarefas {
// 		if tarefa.ID == t.ID {
// 			models.Tarefas = append(models.Tarefas[:i], models.Tarefas[i+1:]...)
// 			break

// 		}
// 	}
// }

// func ModificarTarefa(w http.ResponseWriter, r *http.Request) {
// 	var t models.Tarefa

// }
