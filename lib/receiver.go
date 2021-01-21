package lib

import (
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
)

const tokenvals = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const tokenlen = 5

func tokenGenerator() string {
	tokenlistlen := len(tokenvals)
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, tokenlen)
	for i := range b {
		b[i] = tokenvals[rand.Intn(tokenlistlen)]
	}
	return string(b)
}

// Receiver object
type Receiver struct {
	ws    *websocket.Conn
	token string
}

// register the receiver into the db
func (r *Receiver) register() {
	GDB.Set(r.token, &r, cache.DefaultExpiration)
}

// unregister the receiver
func (r *Receiver) unregister() {
	GDB.Delete(r.token)
}

// SendSecret to the receiver
func (r *Receiver) SendSecret(s Secret) {
	r.ws.WriteJSON(s)
	r.unregister()
}

// NewReceiver handles the creation of a new receiver
func NewReceiver(w *websocket.Conn) *Receiver {
	// create unique token
	var token string
	for {
		token = tokenGenerator()
		_, exists := GDB.Get(token)
		if !exists {
			break
		}
	}
	// insert into the db and return
	r := Receiver{w, token}
	go r.register()
	return &r
}
