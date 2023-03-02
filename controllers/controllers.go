package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TiMattos/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page 2")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)

}
