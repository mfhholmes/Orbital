package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)
//import("time")


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
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("public_html")))
  http.Handle("/", r)

  fmt.Println("firing up webserver")
  go http.ListenAndServe(":8080", nil)

  fmt.Println("firing up actionList")
  actions := new(ActionList)
  controlChan := make(chan controlInstruction)
  actions.mainLoop(controlChan)

  fmt.Println("server ended")
}

