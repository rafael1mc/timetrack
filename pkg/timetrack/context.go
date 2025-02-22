package timetrack

import (
	"context"
)

type timeNodeKey struct{}

// BranchFrom will retrieve a timeNode from ctx and create a new child counter called name
// If there's no node in context, it will create a new one and put it inside returned ctx
func BranchFrom(ctx context.Context, name string) (context.Context, *TimeNode) {
	var newNode *TimeNode
	if node, ok := ctx.Value(timeNodeKey{}).(*TimeNode); ok {
		newNode = node.Branch(name)
	} else {
		newNode = NewNode(name)
	}
	newCtx := WithTimeNode(ctx, newNode)
	return newCtx, newNode
}

// WithtimeNode returns a new context with a timeNode inside
func WithTimeNode(ctx context.Context, node *TimeNode) context.Context {
	return context.WithValue(ctx, timeNodeKey{}, node)
}
