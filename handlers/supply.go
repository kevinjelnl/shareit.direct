package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Supply handles the supplier role by sending the secret to the receiver
func Supply(c echo.Context) error {
	return c.String(http.StatusOK, "Supply")
}
