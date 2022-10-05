package main

import (
	"minijo/db"
	"minijo/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	DB := db.Init()

	routes.UserRoutes(e, DB)
	routes.KeyboardRoutes(e, DB)
	routes.ReviewsRoutes(e, DB)

	e.Logger.Fatal(e.Start(":5000"))
}
