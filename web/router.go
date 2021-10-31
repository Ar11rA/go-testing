package web

import (
	"basics/web/handlers"
	"basics/web/services"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	helloHandler := handlers.NewHelloHander(services.NewHelloService())
	router.HandleFunc("/hello", helloHandler.GetHello).Methods("GET")
	return router
}
