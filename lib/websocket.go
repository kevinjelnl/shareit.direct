package lib

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (r Receiver) listen() {
	// close the socket on close
	defer func() {
		r.ws.Close()
		r.unregister()
	}()
	// loop the connection
	for {
		_, _, err := r.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

// WebsocketHandler handles the websocket connection for /ws
func WebsocketHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	header := w.Header()
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "*")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
	}
	receiver := NewReceiver(ws)
	receiver.ws.WriteMessage(websocket.TextMessage, []byte(receiver.token))
	go receiver.listen()
	fmt.Println("supplied token")
}
