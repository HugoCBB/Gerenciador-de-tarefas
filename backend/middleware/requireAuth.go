package middleware

import (
	"net/http"
	"time"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/config"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/database"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Responsavel por autenticar usuario e protejer a rota
func RequireAuth(c *gin.Context) {
	// Pega o cookie salvo

	tokenKey := config.NewTokenConfig()
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey.Secret_key), nil

	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	// Checa o tempo de expiracao do token, se for menor que o tempo estibulado ele nao autoriza
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Verifica se o usuario existe
		var user models.User
		database.DB.First(&user, claims["sub"])
		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Com todas as  verificacoes feitas, e setado a chave de user
		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
