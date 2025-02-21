package main

import (
	"context"
	"time"

	"github.com/rafael1mc/timetrack/pkg/timetrack"
)

func repositoryFunc(ctx context.Context) {
	_, timer := timetrack.BranchFrom(ctx, "repositoryFunc")

	time.Sleep(100 * time.Millisecond)

	timer.Stop()
}
