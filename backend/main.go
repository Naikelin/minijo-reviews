package main

import (
    "backend/configs"
    "github.com/labstack/echo/v4"
    "backend/routes"
)

func main() {
    e := echo.New()

    configs.ConnectDB()
    routes.KbRoute(e)

    e.Logger.Fatal(e.Start(":6000"))
}
