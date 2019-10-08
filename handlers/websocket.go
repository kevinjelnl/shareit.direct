package handlers

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/kevinjelnl/shareit.direct/models"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Websocket handles the websocket connection
func Websocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		ws.Close()
	}

	// remove
	_, msg, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
	}
	if err := json.Unmarshal([]byte(msg)); err != nil {
		log.Fatal(err)
	}
	print()

	receiver := models.RegisterReceiver(ws)
	w := models.WsMessage{Status: "init",
		TokenStr: receiver.Token.Uuid,
	}
	ws.WriteJSON(w)
	return nil
}
