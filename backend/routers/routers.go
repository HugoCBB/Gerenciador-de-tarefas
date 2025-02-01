package routers

import (
	"log"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/controllers"
	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/middleware"
	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})

}

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(enableCORS, middleware.ContentTypeMiddleware)
	r.HandleFunc("/tarefa", controllers.Tarefas).Methods("GET")
	r.HandleFunc("/tarefa/{id}", controllers.ObterTarefa).Methods("GET")
	r.HandleFunc("/tarefa/adicionar", controllers.AdicionarTarefa).Methods("POST")
	r.HandleFunc("/tarefa/deletar/{id}", controllers.DeletarTarefa).Methods("DELETE")
	r.HandleFunc("/tarefa/modificar/{id}", controllers.ModificarTarefa).Methods("PUT")

	r.HandleFunc("/user", controllers.Usuarios).Methods("GET")
	r.HandleFunc("/user/cadastrar", controllers.CadastrarUsuario).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
