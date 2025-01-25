package routers

import (
	"log"
	"net/http"

	"github.com/HugoCBB/Gerenciador-de-tarefas/backend/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Tarefas).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
