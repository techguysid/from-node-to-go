package main

import (
	"net/http"
)

func indexHandler(res http.ResponseWriter, r *http.Request) {
	res.Write([]byte("Hello World"))
}

func pathAHandler(res http.ResponseWriter, r *http.Request) {
	res.Write([]byte("Yoi've hit /a path"))
}

func main() {

	//Creating your own multiplexer and telling it what handler to invoke when based on url
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/a", pathAHandler)

	// Here we pass our own mux which we created instead of nil
	http.ListenAndServe(":8084", mux)

}
