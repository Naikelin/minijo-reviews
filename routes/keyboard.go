package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"minijo/models"
)

func KeyboardRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/keyboards", func(c echo.Context) error {
		keyboards := []models.Keyboard{}
		result := db.Find(&keyboards)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, keyboards)
	})

	e.POST("/keyboards", func(c echo.Context) error {
		keyboards := []models.Keyboard{}
		if err := c.Bind(&keyboards); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := db.Create(&keyboards)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Keyboard created")
	})
}
