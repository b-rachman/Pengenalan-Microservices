package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/b-rachman/Pengenalan-Microservices/auth-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/validate-admin", http.HandlerFunc(handler.ValidateAdmin))
	fmt.Println("Auth ser vice listen on 8001")
	log.Panic(http.ListenAndServe(":8001", router))

}
