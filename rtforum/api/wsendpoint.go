package api

import (
	"encoding/json"
	"forum/config"
	"forum/utils"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var NC *ConnectionMap

func init() {
	NC = NewConnectionMap()
}

type ConnectionMap struct {
	connections map[int][]*websocket.Conn // Store a slice of connections per user ID
	mu          sync.Mutex
}

func NewConnectionMap() *ConnectionMap {
	return &ConnectionMap{
		connections: make(map[int][]*websocket.Conn),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (NC *ConnectionMap) AddConnection(userID int, conn *websocket.Conn) {
	NC.mu.Lock()
	defer NC.mu.Unlock()
	NC.connections[userID] = append(NC.connections[userID], conn)
}

func (NC *ConnectionMap) RemoveConnection(userID int, conn *websocket.Conn) {
	NC.mu.Lock()
	defer NC.mu.Unlock()
	if conns, ok := NC.connections[userID]; ok {
		for i, c := range conns {
			if c == conn {
				NC.connections[userID] = append(conns[:i], conns[i+1:]...)
				break
			}
		}
		if len(NC.connections[userID]) == 0 {
			delete(NC.connections, userID)
		}
	}
}

func (NC *ConnectionMap) GetConnections(userID int) ([]*websocket.Conn, bool){
	NC.mu.Lock()
	defer NC.mu.Unlock()
	conns, ok := NC.connections[userID]
	return conns, ok
}

func (NC *ConnectionMap) BroadcastMessage(userID int, msg []byte , messagetype int) error {
	NC.mu.Lock()
	defer NC.mu.Unlock()
	if connections, exists := NC.connections[userID]; exists {
		for _, conn := range connections {
			if err := conn.WriteMessage(messagetype, msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	sessionID := utils.GetSessionCookie(r)
	session, err := config.SESSION.GetSession(sessionID)
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	NC.AddConnection(int(session.UserId), conn)
	for {
		messagetype, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var message struct {
			Type string `json:"type"`
			Receiver int `json:"receiver"`
			Sender string `json:"sender"`
			SenderID int `json:"sender_id"`
			Data string `json:"data"`
		}
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println(err)
			continue
		}
		if message.Type == "typing" {
			message.Sender = session.Username
			message.SenderID = int(session.UserId)
			typingMessage, _ := json.Marshal(message)
			if err := NC.BroadcastMessage(message.Receiver, typingMessage, messagetype); err != nil {
				log.Println(err)
			}
			continue
		} else {
			exists := false
			err = 
		}

	}
}