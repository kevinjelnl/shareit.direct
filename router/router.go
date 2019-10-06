package router

import (
	"github.com/foolin/goview/supports/echoview"
	"github.com/kevinjelnl/shareit.direct/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New creates a fresh Echo instance
func New() *echo.Echo {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// middleware for serving templates and static files
	e.Renderer = echoview.Default()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "./views/src",
		Browse: false,
	}))
	e.File("/favicon.ico", "./views/src/favicon.ico")

	// endpoints
	e.GET("/", handlers.Index)
	e.GET("/ws", handlers.Websocket)
	e.POST("/supply", handlers.Supply)
	return e
}
