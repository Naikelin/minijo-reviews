package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"minijo/models"
)

func UserRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/login", func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		findUser := models.User{}

		result := db.Where("username = ?", user.Username).First(&findUser)

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(result.Error)
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		if (findUser.Username == user.Username) && (findUser.Password == user.Password) {
			return c.JSON(http.StatusOK, "login successful")
		} else {
			return c.JSON(http.StatusUnauthorized, "No login 4 u >:c")
		}
	})

	e.POST("/register", func(c echo.Context) error {
		user := models.User{}
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		result := db.Create(&user)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, result.Error)
		}

		return c.JSON(http.StatusOK, "User created")
	})
}
