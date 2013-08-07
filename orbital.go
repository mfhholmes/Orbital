package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//import("time")
type Settings struct{
  Port string
}

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

func buildHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/population", pop).Methods("GET")
	r.HandleFunc("/population", setPop).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public_html")))
	return r
}

func main() {

  //TODO: argument might specify config file/settings
  var defaultSettings = Settings{Port:":8080"}
  
	fmt.Println("firing up actionList")
	actions := new(ActionList)
	controlChan := make(chan controlInstruction)
	go actions.mainLoop(controlChan)

	fmt.Println("firing up webserver")
	handler := buildHandler()

	http.ListenAndServe(defaultSettings.Port, handler)

	fmt.Println("server ended")
}
