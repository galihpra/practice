package auth

import (
	"day-22/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) LoginHandler(c echo.Context) error {
	var currentUser = new(model.User)
	if err := c.Bind(currentUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	qry := as.DB.Where("username = ? AND password = ?", currentUser.Username, currentUser.Password).Take(currentUser)
	err := qry.Error

	if err != nil {
		fmt.Println("login process error:", err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "Password atau username salah",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, currentUser)
}
