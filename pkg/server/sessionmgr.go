package server

import (
	"coralgateway/pkg/config"
	"github.com/gorilla/websocket"
	"github.com/speps/go-hashids"
	"log"
	"math/rand"
	"sync"
	"time"
)

const(
	SESSION_STARTED = "STARTED"
	SESSION_INPROGRESS = "INPROGRESS"
	SESSION_ENDED = "ENDED"

)

var store *SessionStore
var sessionMutex = sync.RWMutex{}
func init() {
	tmp := SessionStore{
		sessions:             map[string]*Session{},
	}
	store = &tmp
}

// Check if session is present in session store
func IsSessionExist(sessionId string) bool {
	sessionMutex.RLock()
	defer sessionMutex.RUnlock()
	if _, ok:= store.sessions[sessionId]; ok {
		return true
	}
	return false
}

func GetSessionStore() *SessionStore {
	return store
}

// Get all sessions.
func (store *SessionStore) GetAllSessions() map[string]*Session {
	return store.sessions
}

func (store *SessionStore) GetSession(sessionId string) *Session {
	sessionMutex.RLock()
	defer sessionMutex.RUnlock()
	return store.sessions[sessionId]
}


// Creates a new session associated with a given connection
func CreateNewSession(conn *websocket.Conn, tag string) *Session {
	//log.Println("creating a new session...")
	id := newHashId()
	userConnect := Connection{id, conn.RemoteAddr().String(), conn}
	conngroup := ConnectionGroup{userConnect}
	creationTime := time.Now().Unix()
	session := Session { id, "",conngroup , "STARTED", 	tag,creationTime, creationTime	}
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	store.sessions[id] = &session
	return &session

}

// write the binary data to the socket
func (s *Session) WriteBinary(data []byte) {
	log.Println("Start Writing to the connection")
	s.ConnGroup.UserConnection.Conn.WriteMessage(2, data)
	log.Println("finished writing to the connection")

}

// Write the text data to the socket
func (s *Session) WriteText(data []byte) {
	s.ConnGroup.UserConnection.Conn.WriteMessage(1, data)
	log.Println("finished writing to the connection")
}


// Cleanup work to remove stale sessions which runs every 5 mins.
func CleanupWorker() {
	for  {
		for sessionId, session := range store.GetAllSessions() {
			timeDiff := time.Now().Unix() - session.LastHeartbeatTime
			if timeDiff > config.GetSessionTimeout() {
				delete(store.sessions, sessionId)
			}
		}
		time.Sleep(300*time.Second)
	}
}

func newHashId() string {
	var hd = hashids.NewData()
	hd.Salt = "Coral Server"
	h, err := hashids.NewWithData(hd)
	handleError(err)
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	rand.Seed(time.Now().UnixNano())
	randomness := rand.Int()
	a := []int {year, month, day, hour, minute, second, randomness}
	for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	id, _ := h.Encode(a)
	return id
}

func handleError(err error) {
	if err != nil {
		log.Println("handling error::::", err)

	}
}

// Stores the live session which are currently running in the server
type SessionStore struct {
	sessions map[string]*Session
}

// Session is started when first user connects to it the server
// a unique session Id is given to it.
type Session struct {
	SessionId string
	AuthToken string
	ConnGroup ConnectionGroup
	State string
	Tag string
	CreationTime int64
	LastHeartbeatTime int64
}

// Holds the socket connection and a unique id for it.
type Connection struct {
	Id string
	ClientAddr string
	Conn *websocket.Conn
}

// Connetion Group is to hold multiple connection
type ConnectionGroup struct {
	UserConnection Connection
}