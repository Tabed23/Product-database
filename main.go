package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"./api"
)

func main(){

	r :=  mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", api.GetProducts).Methods("GET")
	r.HandleFunc("/{id}", api.GetProduct).Methods("GET")
	r.HandleFunc("/{id}/{name}/{price}/{quantity}", api.CreateProduct).Methods("POST")
	r.HandleFunc("/{id}", api.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}