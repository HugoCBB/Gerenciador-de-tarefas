package routers

import (
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/controllers"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ContentTypeMiddleware())

	api := r.Group("/api")
	{
		tarefas := api.Group("/tarefas")
		{
			tarefas.GET("/", controllers.Tarefas)
			tarefas.GET("/:id", controllers.ObterTarefas)
			tarefas.POST("/adicionar", controllers.AdicionarTarefa)
			tarefas.DELETE("/deletar/:id", controllers.DeletarTarefa)
			tarefas.PUT("/editar/:id", controllers.EditarTarefa)

		}
		user := api.Group("/user")
		{
			user.GET("/", controllers.Usuarios)
			user.POST("/cadastrar", controllers.Cadastrar)

		}
	}

	r.Run(":8000")
}
