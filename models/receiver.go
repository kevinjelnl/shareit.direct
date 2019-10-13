package models

import "github.com/gorilla/websocket"

// Receiver gets the secret from the supplier
type Receiver struct {
	*Token
	ws *websocket.Conn
}

// RegisterReceiver creates a new receiver
func RegisterReceiver(ws *websocket.Conn) Receiver {
	token := NewToken()
	r := Receiver{token, ws}
	token.registerToken(r)
	return r
}
