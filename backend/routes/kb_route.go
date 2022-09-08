package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func KbRoute(e *echo.Echo) {
	e.POST("/kb", controllers.CreateKb)
	e.GET("/kb/:kbId", controllers.GetAKb)
	e.PUT("/kb/:kbId", controllers.EditAKb)
	e.DELETE("/kb/:kbId", controllers.DeleteAKb)
	e.GET("/kbs", controllers.GetAllKbs)
}
