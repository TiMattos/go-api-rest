package routes

import (
	"log"
	"net/http"

	"github.com/TiMattos/go-api-rest/controllers"
	"github.com/TiMattos/go-api-rest/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	//mux
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AlterarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
