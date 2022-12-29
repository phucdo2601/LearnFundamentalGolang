package controllers

import (
	"log"
	"net/http"
	services "services"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleRequest() {
	InitializeRoutes()
}

func InitializeRoutes() {
	InitRoutesByGorillaMux()
}

func InitRoutesByGorillaMux() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/migration", services.CreateDatabaseSchema).Methods("POST")
	myRouter.HandleFunc("/user", services.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/update/user", services.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/users/getListUser", services.GetAllUsers).Methods("GET")
	myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}
