package handlers

import (
	"net/http"

	"github.com/kevinjelnl/shareit.direct/models"
	"github.com/labstack/echo"
)

// Supply handles the supplier role by sending the secret to the receiver
func Supply(c echo.Context) error {
	s := new(models.WsMessage)
	err := c.Bind(s)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid")
	}
	// transfer
	err = s.TransferSecret(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "could not transfer")
	}
	return c.String(http.StatusOK, "Secret supplied!")
}
