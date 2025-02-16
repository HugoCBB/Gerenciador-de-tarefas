package controllers

import (
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gin-gonic/gin"
)

func Usuarios(c *gin.Context) {
	var u []models.User
	database.DB.Find(&u)
	c.JSON(http.StatusOK, u)
}

func CadastrarUsuario(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.ValidateRequiredFields(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&u)
	c.JSON(http.StatusOK, u)
}
