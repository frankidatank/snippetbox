package main

import (
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	//servemux "/" is a real all we use this to bypass that
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {
	//router or servemux in go terms
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/create", snippetCreate)

	mux.HandleFunc("/snippet/view", snippetView)

	log.Println("Starting server on :4000")
	//server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
