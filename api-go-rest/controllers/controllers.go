package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jclaudiotomasjr/alura/api-go-rest/database"
	"github.com/jclaudiotomasjr/alura/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func AllTech(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "applitcation/json")
	var t []models.Tecnologias
	database.DB.Find(&t)
	json.NewEncoder(w).Encode(t)
}

func ReturnTech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var t []models.Tecnologias
	database.DB.First(&t, id)
	json.NewEncoder(w).Encode(t)
}

func CreateNewTech(w http.ResponseWriter, r *http.Request) {
	var newTech models.Tecnologias
	json.NewDecoder(r.Body).Decode(&newTech)
	database.DB.Create(&newTech)
	json.NewEncoder(w).Encode(newTech)
}

func DeleteATech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var t models.Tecnologias
	database.DB.Delete(&t, id)
	json.NewEncoder(w).Encode("Tecnologia deletada com sucesso!")
}

func UpdateATech(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var t models.Tecnologias
	database.DB.First(&t, id)
	json.NewDecoder(r.Body).Decode(&t)
	database.DB.Save(&t)
	json.NewEncoder(w).Encode("Tecnologia alterada com sucesso!")
}
