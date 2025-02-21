package main

import (
	"fmt"
	"net/http"

	"github.com/rafael1mc/timetrack/pkg/timetrack"
)

func measuring(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := timetrack.NewNode("middleware")
		ctx := timetrack.WithTimeNode(r.Context(), timer)

		f(w, r.WithContext(ctx))

		fmt.Println(timer)
	}
}
