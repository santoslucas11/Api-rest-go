package main

import (
	models "api/Models"
	routes "api/Routes"
	"api/database"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConnectDB()

	fmt.Println("server on")
	routes.HandleRequest()
}
