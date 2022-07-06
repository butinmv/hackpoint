package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var teams = []Team{
	{1, "3CRABS"},
	{2, "CRABS"},
	{3, "C"},
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/teams", func(c echo.Context) error {
		return c.JSON(http.StatusOK, teams)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
