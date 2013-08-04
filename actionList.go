package main

import("sort")
import("time")
import("fmt")


var isRunning bool

type processFunc func()

type controlInstruction struct{
  runLoop bool
  delay time.Duration
}

type actionItem struct{
  execTime time.Time
  process processFunc
}

//ActionList is a list of Action Items, which is then
type ActionList struct{
  items []*actionItem
}
func (al *ActionList) Pop() *actionItem{
  if (len(al.items) == 0){
    // return an item that does nothing and is set to run an hour from now
    return &actionItem{execTime:time.Now().Add(time.Hour),process:func(){}}
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
    // return an item that does nothing and is set to run a year from now
    return &actionItem{execTime:time.Now().Add(time.Hour),process:func(){}}
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
func (al ActionList) listStatus() string {
  if(len(al.items) > 0){
    return fmt.Sprintf("items: %d earliest: %s",len(al.items),al.items[0].execTime.String())
  }
  return "no items in queue"
}

func (al *ActionList) mainLoop(controlChannel chan controlInstruction){
  timeout := make(chan bool, 1)
  runLoop := true
  delay := (100 * time.Millisecond)
  for runLoop{
    isRunning = true
    if(al.Len() > 0){
      for al.Peek(0).execTime.Before(time.Now()){
        go al.Pop().process()
      }
    }
    // kicks off a 100ms delay, which is used to check the control channel for control instructions
    go func() {
      time.Sleep(delay)
      timeout <- true
    }()
    select {
      case command := <-controlChannel:
        runLoop = command.runLoop
        delay = command.delay
      case <-timeout:
        // timeout, so continue around the loop
    }
  }
  isRunning = false
}