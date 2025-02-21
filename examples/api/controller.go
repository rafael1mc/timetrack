package main

import (
	"net/http"

	"github.com/rafael1mc/timetrack/pkg/timetrack"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	ctx, timer := timetrack.BranchFrom(r.Context(), "controller")
	defer timer.Stop()

	serviceFunc1(ctx)
	serviceFunc2(ctx)
	serviceFunc3(ctx)

	w.Write([]byte("Measurement complete, check terminal..."))
}
