package messaging

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"ride-sharing/shared/contracts"

	"github.com/gorilla/websocket"
)

var (
	ErrConnectionNotFound = errors.New("connection not found")
)

// connWrapper is a wrapper around the websocket connection to allow for thread-safe operations
// This is necessary because the websocket connection is not thread-safe
type connWrapper struct {
	conn  *websocket.Conn
	mutex sync.Mutex
}

type ConnectionManager struct {
	connections map[string]*connWrapper // Local connections storage (userId -> connection)
	mutex       sync.RWMutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for now
	},
}

// Note that on multiple instances of the API gateway, the connection manager needs to store the connections on a separate shared storage.
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]*connWrapper),
	}
}

func (cm *ConnectionManager) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (cm *ConnectionManager) Add(id string, conn *websocket.Conn) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.connections[id] = &connWrapper{
		conn:  conn,
		mutex: sync.Mutex{},
	}

	log.Printf("Added connection for user %s", id)
}

func (cm *ConnectionManager) Remove(id string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.connections, id)
}

func (cm *ConnectionManager) Get(id string) (*websocket.Conn, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	wrapper, exists := cm.connections[id]
	if !exists {
		return nil, false
	}
	return wrapper.conn, true
}

func (cm *ConnectionManager) SendMessage(id string, message contracts.WSMessage) error {
	cm.mutex.RLock()
	wrapper, exists := cm.connections[id]
	cm.mutex.RUnlock()

	if !exists {
		return ErrConnectionNotFound
	}

	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	return wrapper.conn.WriteJSON(message)
}
