package timetrack

import (
	"fmt"
	"strings"
	"time"
)

type SimpleTimeReporter struct{}

func NewSimpleTimeReporter() *SimpleTimeReporter {
	return &SimpleTimeReporter{}
}

func (r SimpleTimeReporter) Report(node *TimeNode) string {
	var sb strings.Builder
	r.buildReport(&sb, node, 0, node.Duration())
	return sb.String()
}

func (r *SimpleTimeReporter) buildReport(sb *strings.Builder, node *TimeNode, level int, total time.Duration) {
	indent := strings.Repeat("  ", level)
	percentage := float64(node.Duration().Milliseconds()) / float64(total.Milliseconds()) * 100

	if level == 0 {
		sb.WriteString(fmt.Sprintf("%s: %v (100%%)\n", node.Name(), node.Duration()))
	} else {
		sb.WriteString(fmt.Sprintf("%s|-- %s: %v (%.1f%%)\n", indent, node.Name(), node.Duration(), percentage))
	}

	for _, child := range node.Children() {
		r.buildReport(sb, child, level+1, total)
	}
}
