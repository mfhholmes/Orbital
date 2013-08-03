package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "time"
  "container/list"
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

func mainLoop(){
  funclist := list.New()
  testfunc := func(){}
  testfunc = func(){
    fmt.Println("test" + time.Now().String())
    funclist.PushBack(testfunc)
  }
  funclist.PushBack(testfunc)
  for{
    runitem := funclist.Front()
    runfunc := runitem.Value
    go runfunc()
    runitem.Next()
    time.Sleep(100 * time.Millisecond)
  }
}
func main() {

  r := mux.NewRouter()
  r.HandleFunc("/population", pop).Methods("GET")
  r.HandleFunc("/population", setPop).Methods("POST")
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("public_html")))
  http.Handle("/", r)
  go http.ListenAndServe(":8080", nil)
  mainLoop();
  fmt.Println("server ended")
}

