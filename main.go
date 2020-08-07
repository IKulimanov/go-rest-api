package main

import (
	"github.com/IKulimanov/go-rest-api/auth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//инициализация роутера
	Myrouter := mux.NewRouter()
	//домашняя страничка
	Myrouter.Handle("/",http.FileServer(http.Dir("./views")))

	Myrouter.HandleFunc("/token", auth.CreateToken).Methods("POST")



	//старт сервера
	log.Fatal(http.ListenAndServe(":8000", Myrouter))
}

var NotImplemented = http.HandlerFunc(func(write http.ResponseWriter, r *http.Request){
	write.Write([]byte("Not Implemented"))
})
