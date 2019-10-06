package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index handles the landing page of shareit.direct
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{})
}
