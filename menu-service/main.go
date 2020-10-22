package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/b-rachman/Pengenalan-Microservices/menu-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu Service Listern to :8000")
	log.Panic(http.ListenAndServe(":8000", router))

}
