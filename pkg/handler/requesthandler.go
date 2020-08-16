package handler

import (
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex

type DimensionConnInput struct {
	Id string
	Name string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle connection")

}

func processMessage( msg []byte) {

}