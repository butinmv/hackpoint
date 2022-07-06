package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/teams", getTeams)

	e.Logger.Fatal(e.Start(":1323"))
}

func getTeams(c echo.Context) error {
	teams := []Team{
		{1, "3CRABS"},
		{2, "CRABS"},
		{3, "C"},
	}
	return c.JSON(http.StatusOK, teams)
}
