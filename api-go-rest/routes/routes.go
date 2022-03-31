package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jclaudiotomasjr/alura/api-go-rest/controllers"
	"github.com/jclaudiotomasjr/alura/api-go-rest/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleare)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/tecnologias", controllers.AllTech).Methods("Get")
	r.HandleFunc("/api/tecnologias/{id}", controllers.ReturnTech).Methods("Get")
	r.HandleFunc("/api/tecnologias", controllers.CreateNewTech).Methods("Post")
	r.HandleFunc("/api/tecnologias/{id}", controllers.DeleteATech).Methods("Delete")
	r.HandleFunc("/api/tecnologias/{id}", controllers.UpdateATech).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
