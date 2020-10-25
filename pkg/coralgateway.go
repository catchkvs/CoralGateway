package main

import (
	"coralgateway/pkg/dbconnector"
	"coralgateway/pkg/handler"
	"flag"
	"log"
	"net/http"
)

func main() {
	dbconnector.InitPebbleDB("dbroot/demo")
	log.Println("Starting the server...")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/sync-db", handler.HandleRequest)
	http.HandleFunc("/", ping)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}