package main

import (
	"minijo/db"
	"minijo/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	DB := db.Init()

	routes.UserRoutes(e, DB)
	routes.KeyboardRoutes(e, DB)
	routes.ReviewsRoutes(e, DB)

	e.Logger.Fatal(e.Start(":5000"))
}
