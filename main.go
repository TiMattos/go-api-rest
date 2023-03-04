package main

import (
	"fmt"

	"github.com/TiMattos/go-api-rest/models"
	"github.com/TiMattos/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Thiago", Historia: "historia do Thiago"},
		{Id: 2, Nome: "Stephanie", Historia: "historia do Stephanie"},
	}

	fmt.Println("Iniciando o servidor Rest em GO")
	routes.HandleRequest()

}
