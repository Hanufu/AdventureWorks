package main

import (
	"fmt"

	"github.com/Hanufu/AdventureWorks/models"
	"github.com/Hanufu/AdventureWorks/routes"
)

func main() {

	models.Products = []models.Product{
		{
			Id:                    1,
			ProductKey:            214,
			ProductSubcategoryKey: 31,
			ProductName:           "Sport-100 Helmet, Red",
			ModelName:             "Sport-100",
			ProductDescription:    "Universal fit, well-vented, lightweight , snap-on visor.",
			ProductColor:          "Red",
			ProductSize:           "XL",
			ProductStyle:          "U",
			ProductCost:           13.0863,
			ProductPrice:          34.99,
		},
		{
			Id:                    2,
			ProductKey:            215,
			ProductSubcategoryKey: 31,
			ProductName:           "Sport-100 Helmet, Black",
			ModelName:             "Sport-100",
			ProductDescription:    "Universal fit, well-vented, lightweight , snap-on visor.",
			ProductColor:          "Black",
			ProductSize:           "XL",
			ProductStyle:          "U",
			ProductCost:           13.0863,
			ProductPrice:          34.99,
		},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
