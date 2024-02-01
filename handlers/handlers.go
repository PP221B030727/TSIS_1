package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Character struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	House    string `json:"house"`
	Blood    string `json:"blood"`
	Birthday string `json:"birthday"`
}

var characters []Character

func init() {
	loadData()
}

func loadData() {
	file, err := os.Open("./data/characters.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&characters)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	for _, character := range characters {
		if strings.ToLower(character.Name) == name {
			json.NewEncoder(w).Encode(character)
			return
		}
	}
	http.Error(w, "Character not found", http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The Harry Potter API is alive! Welcome !!!")
}
