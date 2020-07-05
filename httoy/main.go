package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, `{ "name": "Tanaka" }`)
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	httpServer.Addr = ":55555"
	httpServer.ListenAndServe()
}
