package controllers

import (
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Usuarios(c *gin.Context) {
	var u []models.User
	database.DB.Preload("Tarefas").Find(&u)
	c.JSON(http.StatusOK, u)
}

func Cadastrar(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.ValidateRequiredFields(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Transforma a senha em uma hash
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Nome:  u.Nome,
		Email: u.Email,
		Senha: string(hash),
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func Login() {

}
