package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TiMattos/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page 2")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	achei := false
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			achei = true
			json.NewEncoder(w).Encode(personalidade)
		}

	}
	if !achei {
		fmt.Fprint(w, "Id não localizado")
	}
}
