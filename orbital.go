package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

type Message struct {
  CurrentPopulation int
}

func pop(response http.ResponseWriter, request *http.Request) {
  status := Message{5}
  json, _ := json.Marshal(status)
  fmt.Fprintf(response, "%s", json)
}

func setPop(response http.ResponseWriter, request *http.Request) {
  start := request.FormValue("startPopulation")
  growth := request.FormValue("growthRate")
  fmt.Fprintf(response, "%s + %s per ?", start, growth)
}

func main() {

  r := mux.NewRouter()
  r.HandleFunc("/population", pop).Methods("GET")
  r.HandleFunc("/population", setPop).Methods("POST")
  http.Handle("/population", r)
  http.Handle("/",http.FileServer(http.Dir("public_html")))
  http.ListenAndServe(":8080", nil)
}