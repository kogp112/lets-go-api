package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// struct is a collection fields
type Person struct {
	ID   string `json:"ID"`
	Name string `json:"Title"`
	Age  int    `json:"Age"`
}

// slice doesn't have size, it's dinamically flexible and common
var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if item.ID == params["userID"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Person
	_ = json.NewDecoder(r.Body).Decode(&data)
	data.ID = strconv.Itoa(rand.Intn(1000000))
	persons = append(persons, data)
	json.NewEncoder(w).Encode(&data)
}

func main() {
	// initialize new router
	router := mux.NewRouter().StrictSlash(true)
	// create mock datas
	persons = append(persons, Person{ID: "1", Name: "Kanako", Age: 30})
	persons = append(persons, Person{ID: "2", Name: "Taro", Age: 33})
	// create api endpoints
	router.HandleFunc("/events", getPersons).Methods(http.MethodGet)
	router.HandleFunc("/events/{userID}", getPerson).Methods(http.MethodGet)
	router.HandleFunc("/events", createPersons).Methods(http.MethodPost)

	http.ListenAndServe(":8081", router)
}
