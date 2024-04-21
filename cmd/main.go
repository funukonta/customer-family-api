package main

import (
	"customer-data-api/internal/handlers"
	"customer-data-api/internal/repositories"
	"customer-data-api/internal/services"
	"customer-data-api/pkg"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	db := pkg.ConnectPostgres()

	repo := repositories.NewCustomerRepo(db)
	service := services.NewCustomerService(repo)
	handler := handlers.NewCustomerHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/customers", handler.Create).Methods("POST")
	r.HandleFunc("/customers", handler.GetAllCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", handler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", handler.DeleteCustomer).Methods("DELETE")

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Println("api server listening to port", port)

	http.ListenAndServe(port, r)
}
