package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"minijo/models"
)

func avrgStars(keyboardID int, db *gorm.DB) float64 {
	reviews := []models.Review{}
	db.Where("keyboard_id = ?", keyboardID).Find(&reviews)

	var sum int
	for _, review := range reviews {
		sum += review.Stars
	}

	return float64(sum) / float64(len(reviews))
}

func KeyboardRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/keyboards", func(c echo.Context) error {
		keyboards := []models.Keyboard{}
		result := db.Find(&keyboards)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		for i, keyboard := range keyboards {
			keyboards[i].Stars = avrgStars(keyboard.ID, db)
		}

		return c.JSON(http.StatusOK, keyboards)
	})

	e.POST("/keyboards", func(c echo.Context) error {
		keyboards := models.Keyboard{}
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
