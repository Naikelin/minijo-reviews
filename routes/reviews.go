package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"minijo/models"
)

func ReviewsRoutes(e *echo.Echo, db *gorm.DB) {

	/*
		Obtiene todas las reviews de la base de datos
		Ruta: {GET}
			$URL/getReviews
	*/

	e.GET("/getReviews", func(c echo.Context) error {
		reviews := []models.Review{}
		result := db.Find(&reviews)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, reviews)
	})

	/*
		Obtiene una review a partir de su id
		Ruta: {GET}
			$URL/getReview/:id
	*/

	e.GET("/getReview/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		result := db.First(&models.Review{}, id)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}
		return c.JSON(http.StatusOK, result)
	})

	/*
		Crea una review. Necesita:
			- id del usuario {user_id} // NOT NULL
			- id del teclado {keyboard_id} // NOT NULL
			- descripción {description} // default ''
			- stars {stars} // default 0

		EJ:
			{
				"keyboard_id": 2,
				"user_id": 1,
				"description": "este teclado ta bonito",
				"stars": 4
			}
		Ruta: {POST}
			$URL/createReview
	*/

	e.POST("/createReview", func(c echo.Context) error {
		review := models.Review{}

		if err := c.Bind(&review); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := db.Create(&review)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Review created")
	})

	/*
		Edita una review. Necesita:
			- id de la review {id} // NOT NULL
			- descripción {description} // Enviar valor previo, por cualquier problema
			- stars {stars} // Enviar valor previo, por cualquier problema

		EJ:
			{
				"id": 1,
				"stars": 3,
				"description": "creo q ya no es tan sex0"
			}
		Ruta: {PUT}
			$URL/editReview
	*/

	e.PUT("/editReview", func(c echo.Context) error {
		review := models.Review{}

		if err := c.Bind(&review); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := db.Model(&review).Updates(map[string]interface{}{"stars": review.Stars, "description": review.Description})

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Review updated")
	})

	/*
		Elimina una review. Necesita:
			- id de la review {id} // NOT NULL

		EJ:
			{
				"id": 1,
			}
		Ruta: {DELETE}
			$URL/deleteReview
	*/

	e.DELETE("/deleteReview", func(c echo.Context) error {
		review := models.Review{}

		if err := c.Bind(&review); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := db.Delete(&models.Review{}, review.ID)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "Review deleted")
	})
}
