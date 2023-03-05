package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TiMattos/go-api-rest/database"
	"github.com/TiMattos/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page 2")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	if personalidade.Nome != "" {
		json.NewEncoder(w).Encode(personalidade)
	} else {
		fmt.Fprint(w, "Id não localizado")
	}

}
