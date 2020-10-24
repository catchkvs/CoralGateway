package main

import (
	"coralgateway/pkg/dbconnector"
	"coralgateway/pkg/handler"
	"flag"
	"log"
	"net/http"
)

func main() {
	dbconnector.Init()
	log.Println("Starting the server...")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/sync-pos", handler.HandleRequest)
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