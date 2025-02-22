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
	timer.SetReporter(CustomReporter{})

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
	/*
		[22:41:40.787] 'root node': 377.737667ms
		[22:41:41.088] 'child node': 76.651916ms
		[22:41:41.088] 'grandchild1 node': 50.498459ms
		[22:41:41.139] 'grandchild2 node': 26.038625ms
	*/
}

// CustomReporter will just output the start time, name and duration
type CustomReporter struct{}

func (cr CustomReporter) Report(node *timetrack.TimeNode) string {
	if len(node.Children()) == 0 {
		return cr.output(node)
	}

	str := cr.output(node)
	for _, c := range node.Children() {
		str += cr.Report(c)
	}

	return str
}

func (cr CustomReporter) output(node *timetrack.TimeNode) string {
	return fmt.Sprintf(
		"[%s] '%s': %v\n",
		node.StartAt().Format("15:04:05.000"),
		node.Name(),
		node.Duration(),
	)
}
