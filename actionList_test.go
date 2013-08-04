package main
import("testing")
import("time")
//import("fmt")

var testCount int

// test action process function
func testRecursiveFunc(t *testing.T, message string, actions *ActionList) processFunc {
  return func(){
    hi := new (actionItem)
    hi.process = testRecursiveFunc(t,message,actions)
    hi.execTime = time.Now().Add(time.Second * 5)
    actions.Push(hi)
    t.Log(message)
    testCount++
  }
}

func testFunc(t *testing.T,message string, actions *ActionList) processFunc {
  return func(){
    t.Log(message)
    testCount++
  }
}
func TestCreate(t *testing.T){
  var actions ActionList
  test := new (actionItem)
  test.process = testFunc(t,"this should never fire",&actions)
  test.execTime = time.Now().Add(time.Second * 1)
  actions.Push(test)
  if(actions.Len() < 1){
    t.Fatal("Creating item failed to increase Len")
  }

  test = new (actionItem)
  test.process = testFunc(t,"this should never fire",&actions)
  test.execTime = time.Now().Add(time.Second * 2)
  actions.Push(test)
  if(actions.Len() < 2){
    t.Fatal("Creating item failed to increase Len")
  }
}

func TestRun(t *testing.T){
  var actions ActionList
  test := new (actionItem)
  test.process = testFunc(t,"this should never fire",&actions)
  test.execTime = time.Now().Add(time.Second * 1)
  actions.Push(test)
  t.Log("created first TestRun item")
  if(actions.Len() < 1){
    t.Fatal("Creating item failed to increase Len")
  }
  testCount = 0
  controlChan := make(chan controlInstruction)
  t.Log("starting TestRun loop")
  go actions.mainLoop(controlChan)
  time.Sleep(time.Second*3)
  if(testCount == 0){
    t.Fatal("test item didn't run")
  }
  instr := controlInstruction{runLoop:false,delay:time.Second}
  controlChan <- instr
}
func TestRunControls(t *testing.T){
  var actions ActionList
  controlChan := make(chan controlInstruction)
  go actions.mainLoop(controlChan)

  time.Sleep(time.Second*1)
  instr := controlInstruction{runLoop:false,delay:time.Second}
  controlChan <- instr
  time.Sleep(time.Second*1)
  if(isRunning){
    t.Fatal("switching runLoop didn't stop it running")
  }
}