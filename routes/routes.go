package routes

import (
	"log"
	"net/http"

	"github.com/Hanufu/AdventureWorks/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/products", controllers.Produts).Methods("GET")
	r.HandleFunc("/products/{productKey}", controllers.ProductId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
