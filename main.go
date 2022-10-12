package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	LoadAppConfig()
	Connect(AppConfig.ConnectionString)
	Migrate()
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user", CreateUsers).Methods("POST")
	router.HandleFunc("/", Health).Methods("GET")
}
