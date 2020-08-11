package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("input")
	if err != nil {
		fmt.Println("set a new cookie.")

		// Cookieにデータがセットされていない状態でリクエストがきた場合はデフォルトの値をSetする。
		value := "default value"
		cookie := new(http.Cookie)
		cookie.Name = "input"
		cookie.Value = value
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Path = "/"
		c.SetCookie(cookie)

		return c.String(http.StatusCreated, "set a new cookie: "+value)
	}
	fmt.Println("------------- cookie content -----------------")
	fmt.Println("cookie name: " + cookie.Name)
	fmt.Println("cookie value: " + cookie.Value)
	fmt.Println("------------- cookie content -----------------")
	return c.String(http.StatusOK, "read a cookie: "+cookie.Value)
}
