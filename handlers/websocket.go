package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Websocket handles the websocket connection
func Websocket(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
