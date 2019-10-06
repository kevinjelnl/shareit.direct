package handlers

import (
	"log"
	"net/http"

	"github.com/kevinjelnl/shareit.direct/models"
	"github.com/labstack/echo"
)

// Index handles the landing page of shareit.direct
func Index(c echo.Context) error {
	log.Print(models.NewToken())
	return c.Render(http.StatusOK, "index", echo.Map{})
}
