package main

import (
	"fmt"
	"time"

	"github.com/rafael1mc/timetrack/pkg/timetrack"
)

func main() {
	exampleFunction()
}

func exampleFunction() {
	timer := timetrack.NewNode("root node")
	time.Sleep(300 * time.Millisecond)

	childExecution := timer.Branch("child node")

	grandChildExecution1 := childExecution.Branch("grandchild1 node")
	time.Sleep(50 * time.Millisecond)
	grandChildExecution1.Stop()

	grandChildExecution2 := childExecution.Branch("grandchild2 node")
	time.Sleep(25 * time.Millisecond)
	grandChildExecution2.Stop()

	childExecution.Stop()

	timer.Stop()

	fmt.Println(timer)
	fmt.Println(childExecution)
	// fmt.Println(grandChildExecution1)
	// fmt.Println(grandChildExecution2)
}
