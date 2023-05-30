package main

import (
	"fmt"

	"github.com/Hanufu/AdventureWorks/database"
	"github.com/Hanufu/AdventureWorks/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
