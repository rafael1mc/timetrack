package timetrack

// type TimeReporter interface {
// 	Report(node *TimeNode) string
// }

// type SimpleTimeReporter struct {
// 	node  *TimeNode
// 	level int
// 	total time.Duration
// }

// func (r SimpleTimeReporter) Report(node *TimeNode) string {
// 	var sb strings.Builder
// 	r.buildReport(&sb, r.node, r.level, r.total)
// 	return sb.String()
// }

// func (r *SimpleTimeReporter) buildReport(sb *strings.Builder, node *TimeNode, level int, total time.Duration) {
// 	indent := strings.Repeat("  ", level)
// 	percentage := float64(node.duration.Milliseconds()) / float64(total.Milliseconds()) * 100

// 	if level == 0 {
// 		sb.WriteString(fmt.Sprintf("%s: %v (100%%)\n", node.name, node.duration))
// 	} else {
// 		sb.WriteString(fmt.Sprintf("%s|-- %s: %v (%.1f%%)\n", indent, node.name, node.duration, percentage))
// 	}

// 	for _, child := range node.children {
// 		r.buildReport(sb, child, level+1, total)
// 		// r.buildReport(sb, child, level+1, node.duration)
// 	}
// }
