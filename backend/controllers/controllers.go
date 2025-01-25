package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
)

func Tarefas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Tarefas)
}
