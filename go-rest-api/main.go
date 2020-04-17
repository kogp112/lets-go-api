package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type event struct {
	ID   string `json:"ID"`
	Name string `json:"Title"`
	Age  int    `json:"Age"`
}

type fetchEventsAll []event

var events = fetchEventsAll{
	{
		ID:   "1",
		Name: "Kana",
		Age:  30,
	},
}

func get(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error happend")
	}
	json.Unmarshal(reqBody, &newEvent)
	json.NewEncoder(w).Encode(events[1])
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := 1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "nothing userID!"}`))
			return
		}
	}
	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "location": "%s" }`, userID, location)))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/lets-go-api/v1").Subrouter()
	router.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/id/{userID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}
