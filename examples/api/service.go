package main

import (
	"context"
	"time"

	"github.com/rafael1mc/timetrack/pkg/timetrack"
)

func serviceFunc1(ctx context.Context) {
	ctx, timer := timetrack.BranchFrom(ctx, "serviceFunc1")
	defer timer.Stop()

	time.Sleep(100 * time.Millisecond)
	repositoryFunc(ctx)
}

func serviceFunc2(ctx context.Context) {
	ctx, timer := timetrack.BranchFrom(ctx, "serviceFunc2")
	defer timer.Stop()

	time.Sleep(200 * time.Millisecond)

	subMeasurement := timer.Branch("serviceFunc2 sub measurement")
	time.Sleep(50 * time.Millisecond)
	subMeasurement.Stop()
}

func serviceFunc3(ctx context.Context) {
	ctx, timer := timetrack.BranchFrom(ctx, "serviceFunc3")

	time.Sleep(300 * time.Millisecond)

	timer.Stop()
}
