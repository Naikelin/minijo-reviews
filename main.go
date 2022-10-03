package main

import (
	"minijo/db"
	"net/http"

	"github.com/labstack/echo"
)

type test struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	db.Init()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, test{Name: "test"})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
