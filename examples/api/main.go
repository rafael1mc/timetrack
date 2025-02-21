package main

import (
	"log"
	"net/http"
)

/*
This example wil simulate measurements across an API call:
endpoint -> middleware -> services -> repository
*/

func main() {
	addr := "localhost:4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", measuring(FooHandler))
	mux.HandleFunc("/favicon.ico", doNothing)

	log.Printf("server is listening at %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

func doNothing(w http.ResponseWriter, r *http.Request) {}
