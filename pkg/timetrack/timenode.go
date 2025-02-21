package timetrack

import (
	"sync"
	"time"
)

type TimeNode struct {
	name      string
	startAt   time.Time
	duration  time.Duration
	children  []*TimeNode
	parent    *TimeNode
	completed bool
	mu        sync.Mutex

	timeProvider TimeProvider
	timeReport   TimeReport
}

func NewNode(name string) *TimeNode {
	initTimeProvider()
	return &TimeNode{
		name:    name,
		startAt: timeProvider.Now(),
	}
}

func (n *TimeNode) Branch(name string) *TimeNode {
	initTimeProvider()
	n.mu.Lock()
	defer n.mu.Unlock()

	child := &TimeNode{
		name:    name,
		startAt: timeProvider.Now(),
		parent:  n,
	}

	n.children = append(n.children, child)
	child.parent = n

	return child
}

func (n *TimeNode) Stop() time.Duration {
	initTimeProvider()
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

func (n *TimeNode) CurrentDuration() time.Duration {
	if !n.completed {
		return timeProvider.Now().Sub(n.startAt)
	}
	return n.duration
}

func (n *TimeNode) String() string {
	return n.Report().Report(n)
}

func (n *TimeNode) Report() TimeReport {
	initTimeProvider()

	if n.timeReport == nil {
		n.timeReport = SimpleTimeReport{
			node:  n,
			total: n.Stop(),
		}
	}

	// TODO this name is weird. TimeReport is not the report, but a reporter(?)
	return n.timeReport
}
