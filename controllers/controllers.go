package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hanufu/AdventureWorks/database"
	"github.com/Hanufu/AdventureWorks/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Products(w http.ResponseWriter, r *http.Request) {
	var p []models.Product
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product

	database.DB.First(&product, id)
	json.NewEncoder(w).Encode(product)
}

func CriarProduto(w http.ResponseWriter, r *http.Request) {

	var novoProduto models.Product
	json.NewDecoder(r.Body).Decode(&novoProduto)
	database.DB.Create(&novoProduto)

	json.NewEncoder(w).Encode(novoProduto)
}

func DeletaProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product

	database.DB.Delete(&product, id)
	json.NewEncoder(w).Encode(product)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var product models.Product

	database.DB.First(&product, id)
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(product)

}
