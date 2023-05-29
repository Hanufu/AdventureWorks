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
	Id := vars["id"]

	for _, product := range models.Products {
		if strconv.Itoa(product.Id) == Id {
			json.NewEncoder(w).Encode(product)
		}
	}
}
