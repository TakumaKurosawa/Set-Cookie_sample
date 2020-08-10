package main

import (
	"set-cookie-sample/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/write/:value", handler.WriteCookie)
	e.GET("/read", handler.ReadCookie)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
