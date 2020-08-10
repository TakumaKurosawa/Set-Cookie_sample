package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func WriteCookie(c echo.Context) error {
	value := c.Param("value")
	cookie := new(http.Cookie)
	cookie.Name = "input"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie: "+value)
}
