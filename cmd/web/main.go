package main

import (
	"log"
	"net/http"
)

func main() {
	//router or servemux in go terms
	mux := http.NewServeMux()
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/create", snippetCreate)

	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting server on :4000")
	//server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
