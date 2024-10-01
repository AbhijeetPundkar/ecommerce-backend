package main

import (
	"github.com/AbhijeetPundkar/ecommerce-backend/internals/db"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main(){
	db.Connect()
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Go Shop"))
	})

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)


	
}