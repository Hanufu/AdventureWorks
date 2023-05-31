package routes

import (
	"log"
	"net/http"

	"github.com/Hanufu/AdventureWorks/controllers"
	"github.com/Hanufu/AdventureWorks/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/products", controllers.Products).Methods("GET")
	r.HandleFunc("/products/", controllers.Products).Methods("GET")

	r.HandleFunc("/products/{id}", controllers.ProductId).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.DeletaProduto).Methods("Delete")
	r.HandleFunc("/products/{id}", controllers.EditarProduto).Methods("Put")

	r.HandleFunc("/products", controllers.CriarProduto).Methods("POST")
	r.HandleFunc("/products/", controllers.CriarProduto).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
