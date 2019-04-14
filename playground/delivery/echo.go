package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello/:username", MainPage())

	e.Start(":8080")
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello World " + username)
	}
}