package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hanufu/AdventureWorks/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Produts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Products)
}
