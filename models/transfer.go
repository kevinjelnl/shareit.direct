package models

import "github.com/labstack/echo"

// WsMessage takes care of sending the message from supplier to receiver
type WsMessage struct {
	Status     string `json:"status"`
	TokenStr   string `json:"token"`
	Secret     string `json:"secret"`
	Passphrase bool   `json:"key"`
}

// TransferSecret ships the secret
func (m *WsMessage) TransferSecret(c echo.Context) error {
	m.Status = "transfer" // set the status
	tokensmap[m.TokenStr].ws.WriteJSON(m)
	return nil
}
