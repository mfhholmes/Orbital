package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "time"
)
import("sort")

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

type processFunc func()

type actionItem struct{
  execTime time.Time
  process processFunc
}

type ActionList struct{
  items []*actionItem
}
func (al *ActionList) Pop() *actionItem{
  if (len(al.items) == 0){
    return nil
  }
  var result = al.items[0]
  al.items = al.items[1:]
  return result
}
func (al ActionList) Peek(index int) *actionItem{
  if(index <0){
    index =0
  }
  if (len(al.items) == 0){
    return nil
  }
  if(index > len(al.items)){
    index = len(al.items)-1
  }
  return al.items[index]
}
func (al *ActionList) Push(newItem *actionItem){
  al.items = append(al.items, newItem)
  sort.Sort(al)
}
func (al ActionList) Len() int {
  return len(al.items)
}

func (al ActionList) Less(i, j int) bool {
  return al.items[i].execTime.Before( al.items[j].execTime)
}
func (al ActionList) Swap(i,j int) {
  al.items[i], al.items[j] = al.items[j], al.items[i]
}
func (al ActionList) printList() string {
  if(len(al.items) > 0){
    return fmt.Sprintf("items: %d earliest: %s",len(al.items),al.items[0].execTime.String())
  }
  return "no items in queue"
}

var actions ActionList

func testFunc(message string) processFunc {
  return func(){
    hi := new (actionItem)
    hi.process = testFunc(message)
    hi.execTime = time.Now().Add(time.Second * 5)
    actions.Push(hi)
    fmt.Println(message)
    fmt.Println(actions.printList())
  }
}
func mainLoop(){
  // init loop - purely for testing
  test := new (actionItem)
  test.process = testFunc("test1 process")
  test.execTime = time.Now().Add(time.Second * 2)
  actions.Push(test)
  fmt.Println(actions.printList())

  time.Sleep(time.Millisecond * 2500)
  test = new (actionItem)
  test.process = testFunc("test2 process")
  test.execTime = time.Now().Add(time.Second * 2)
  actions.Push(test)
  fmt.Println(actions.printList())
  // end of test

  for{
    if(actions.Len() > 0){
      for actions.Peek(0).execTime.Before(time.Now()){
        go actions.Pop().process()
      }
    }
    time.Sleep(time.Millisecond * 100)
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

