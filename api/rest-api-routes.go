package api

import (
	"TSIS_1/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/characters", handlers.GetCharacters).Methods("GET")
	r.HandleFunc("/characters/{name}", handlers.GetCharacter).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
