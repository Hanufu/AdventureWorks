package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hanufu/AdventureWorks/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Produts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Products)
}

func ProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProductKey := vars["productKey"]

	for _, product := range models.Products {
		if strconv.Itoa(product.ProductKey) == ProductKey {
			json.NewEncoder(w).Encode(product)
		}
	}
}
