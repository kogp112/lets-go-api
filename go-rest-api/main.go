package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type event struct {
  ID  string `json:"ID"`
  Name  string `json:"Title"`
  Age  int `json:"Age"`
}

type fetchEventsAll []event

var events = fetchEventsAll{
  {
    ID: "1",
    Name: "Kana",
    Age: 30,
  },
}

func fetch(w http.ResponseWriter, r *http.Request) {
  var newEvent event
  reqBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Fprintf(w, "error happend")
  }
  json.Unmarshal(reqBody, &newEvent)
  events = append(events, newEvent)
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(newEvent)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", fetch)
  log.Fatal(http.ListenAndServe(":8081", router))
}
