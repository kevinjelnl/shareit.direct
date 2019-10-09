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

	receiver := models.RegisterReceiver(ws)
	w := models.WsMessage{Status: "init",
		TokenStr: receiver.Token.Uuid,
	}
	ws.WriteJSON(&w)

	// Remove the token from the tokensmap when the user closes the item
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Print(err)
	}

	// Catch if the user leaves
	var wmsg models.WsMessage
	err = json.Unmarshal(msg, &wmsg)
	if err != nil {
		log.Print(err)
		ws.Close()
	} else {
		if wmsg.Status == "closed" {
			models.RemoveToken(wmsg.TokenStr)
			ws.Close()
		}
	}
	return nil
}
