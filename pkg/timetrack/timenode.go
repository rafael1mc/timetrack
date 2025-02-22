package timetrack

import (
	"sync"
	"time"
)

func NewNode(name string) *TimeNode {
	tp := timeProvider()
	rp := reporter()
	return &TimeNode{
		name:           name,
		startAt:        tp.Now(),
		timeProvider:   tp,
		customReporter: rp,
	}
}

type TimeNode struct {
	name      string
	startAt   time.Time
	duration  time.Duration
	children  []*TimeNode
	parent    *TimeNode
	completed bool
	mu        sync.Mutex

	timeProvider   TimeProvider
	customReporter TimeReporter
}

func (n *TimeNode) Branch(name string) *TimeNode {
	n.mu.Lock()
	defer n.mu.Unlock()

	tp := timeProvider()
	child := &TimeNode{
		name:           name,
		startAt:        tp.Now(),
		parent:         n,
		timeProvider:   tp,
		customReporter: n.customReporter,
	}

	n.children = append(n.children, child)
	child.parent = n

	return child
}

func (n *TimeNode) Stop() time.Duration {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.completed {
		for _, child := range n.children {
			if !child.completed {
				child.Stop()
			}
		}

		n.mu.Unlock()
		d := n.Duration()
		n.mu.Lock()

		n.duration = d
		n.completed = true
	}

	return n.duration
}

func (n *TimeNode) SetReporter(r TimeReporter) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.customReporter = r
}

func (n *TimeNode) Level() int {
	n.mu.Lock()
	if n.parent == nil {
		n.mu.Unlock()
		return 0
	}
	n.mu.Unlock()

	return n.Parent().Level() + 1
}

// access methods

func (n *TimeNode) Name() string {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.name
}

func (n *TimeNode) StartAt() time.Time {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.startAt
}

func (n *TimeNode) Duration() time.Duration {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.completed {
		return n.timeProvider.Now().Sub(n.startAt)
	}

	return n.duration
}

func (n *TimeNode) Children() []*TimeNode {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.children
}

func (n *TimeNode) Parent() *TimeNode {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.parent
}

func (n *TimeNode) IsComplete() bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.completed
}

// overides

func (n *TimeNode) String() string {
	return n.customReporter.Report(n)
}
