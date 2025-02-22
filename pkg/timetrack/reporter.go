package timetrack

type TimeReporter interface {
	Report(node *TimeNode) string
}
