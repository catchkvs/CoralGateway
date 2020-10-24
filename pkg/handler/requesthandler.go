package handler

import (
	"coralgateway/pkg/server"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{}
var mux sync.Mutex

type DimensionConnInput struct {
	Id string
	Name string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle connection")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	session := server.CreateNewSession(c, "Tag1")
	// Send the session id to the client
	msg := server.ServerMsg{server.CMD_ReceiveSessionId, session.SessionId, session.SessionId}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.WriteMessage(1, msgBytes)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {

			log.Println("read:", err)
			break
		}
		log.Printf("message type: %s", mt)
		if mt == 2 {
			log.Println("Cannot process binary message right now")
		} else {
			processMessage(message)
		}
	}
}

func processMessage( msg []byte) {

}