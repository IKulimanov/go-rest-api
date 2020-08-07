package main

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//init router
	Myrouter := mux.NewRouter()

	//Myrouter.HandleFunc("/api/tokens",controller.createTokens).Methods("POST")


	log.Fatal(http.ListenAndServe(":8000", Myrouter))
}
