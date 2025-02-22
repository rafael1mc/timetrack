package timetrack

import (
	"sync"
	"time"
)

func NewNode(name string) *timeNode {
	tp := timeProvider()
	return &timeNode{
		name:         name,
		startAt:      tp.Now(),
		timeProvider: tp,
	}
}

type timeNode struct {
	name      string
	startAt   time.Time
	duration  time.Duration
	children  []*timeNode
	parent    *timeNode
	completed bool
	mu        sync.Mutex

	timeProvider TimeProvider
	// customReporter TimeReporter
}

func (n *timeNode) Branch(name string) *timeNode {
	n.mu.Lock()
	defer n.mu.Unlock()

	tp := timeProvider()
	child := &timeNode{
		name:         name,
		startAt:      tp.Now(),
		parent:       n,
		timeProvider: tp,
	}

	n.children = append(n.children, child)
	child.parent = n

	return child
}

func (n *timeNode) Stop() time.Duration {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.completed {
		for _, child := range n.children {
			if !child.completed {
				child.Stop()
			}
		}

		n.duration = n.CurrentDuration()
		n.completed = true
	}

	return n.duration
}

func (n *timeNode) CurrentDuration() time.Duration {
	if !n.completed {
		return n.timeProvider.Now().Sub(n.startAt)
	}
	return n.duration
}

// func (n *timeNode) String() string {
// 	// TODO init reporter
// 	return n.Report(defaultReporter).Report(n)
// }

// func (n *timeNode) Report(r TimeReporter) {
// 	initTimeProvider()

// 	if n.timeReport == nil {
// 		n.timeReport = SimpleTimeReport{
// 			node:  n,
// 			total: n.Stop(),
// 		}
// 	}

// 	// TODO this name is weird. TimeReport is not the report, but a reporter(?)
// 	return n.timeReport
// }
