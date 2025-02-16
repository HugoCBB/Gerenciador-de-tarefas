package controllers

import (
	"net/http"
	"time"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gin-gonic/gin"
)

func Tarefas(c *gin.Context) {
	var t []models.Tarefa
	database.DB.Find(&t)
	c.JSON(http.StatusOK, t)

}

func ObterTarefas(c *gin.Context) {
	var t models.Tarefa
	id := c.Param("id")

	database.DB.First(&t, id)
	if t.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tarefa nao encontrada"})
		return
	}
	c.JSON(http.StatusOK, t)

}

func AdicionarTarefa(c *gin.Context) {
	var t models.Tarefa
	t.DataCriacao = time.Now()
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := t.ValidateRequiredFields(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&t)
	c.JSON(http.StatusOK, t)
}

func DeletarTarefa(c *gin.Context) {
	var t models.Tarefa
	id := c.Param("id")
	database.DB.Delete(&t, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "tarefa deletada com sucesso"})

}

func EditarTarefa(c *gin.Context) {
	var t models.Tarefa
	id := c.Param("id")
	database.DB.First(&t, id)

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&t).UpdateColumns(t)
	c.JSON(http.StatusOK, t)

}
