package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// struct is a collection fields
type event struct {
	ID   string `json:"ID"`
	Name string `json:"Title"`
	Age  int    `json:"Age"`
}

// slice doesn't have size, it's dinamically flexible and common
type fetchEventsAll []event

var events = fetchEventsAll{
	{
		ID:   "1",
		Name: "Kana",
		Age:  30,
	},
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range events {
		if item.ID == params["userID"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&event{})
}

func main() {
	// initialize new router
	router := mux.NewRouter().StrictSlash(true)
	// create api endpoints
	router.HandleFunc("/events", getEvents).Methods(http.MethodGet)
	router.HandleFunc("/events/{userID}", getEvents).Methods(http.MethodGet)

	http.ListenAndServe(":8081", router)
}
