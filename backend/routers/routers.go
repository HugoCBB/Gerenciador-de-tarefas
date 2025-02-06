package routers

import (
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/controllers"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ContentTypeMiddleware())

	v1 := r.Group("/api/tarefas")
	{
		v1.GET("/", controllers.Tarefas)
		v1.GET("/:id", controllers.ObterTarefas)
		v1.POST("/adicionar", controllers.AdicionarTarefa)
		v1.DELETE("/deletar/:id", controllers.DeletarTarefa)
		v1.PUT("/editar/:id", controllers.EditarTarefa)

	}

	v2 := r.Group("api/user")
	{
		v2.GET("/", controllers.Usuarios)
		v2.POST("/cadastrar", controllers.CadastrarUsuario)
	}
	r.Run(":8000")
}
