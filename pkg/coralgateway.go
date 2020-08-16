package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server...")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", ping)
	log.Fatal(http.ListenAndServe(":4040", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}