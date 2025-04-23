package controllers

import (
	"net/http"
	"time"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/config"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	// verifica a existencia do email no banco
	var existingUser models.User
	if err := database.DB.Where("email = ?", u.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-mail já está em uso"})
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
	c.JSON(http.StatusOK, gin.H{
		"status": "Usuario cadastrado com sucesso",
	})
}

func Login(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica a existencia do email no banco de dados
	user := models.User{}
	if err := database.DB.First(&user, "email = ?", u.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email ou senha inválida"})
		return
	}

	// Compara a senha digitada com o hash de senha salvado no banco
	err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(u.Senha))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Senha invalida"})
		return
	}

	// Gera token jwt
	tokenKey := config.NewTokenConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Assina o token jwt
	tokenString, err := token.SignedString([]byte(tokenKey.Secret_key))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define o cookie com o token jwt
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "Login realizado com sucesso",
		"token":  token,
	})

}

func Validar(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})

}
