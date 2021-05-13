package main

import (
	"fmt"
	"net/http"
)

func main() {

	// handle `/a` route
	http.HandleFunc("/a", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "You've hit /a path")
	})

	// handle `/b` route
	http.HandleFunc("/b", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "You've hit /b path")
	})

	// nil means we are using default server mux `http.DefaultServeMux`
	http.ListenAndServe(":8084", nil)

}
