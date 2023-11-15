package users

import (
	"day-22/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserSystem struct {
	DB *gorm.DB
}

func (us *UserSystem) CreateUser(c echo.Context) error {
	var newUser = new(model.User)
	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	newUser.Status = 2

	qry := us.DB.Create(newUser)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error :", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, newUser)
}

func (us *UserSystem) ReadUser(c echo.Context) error {
	var UserList []model.User

	qry := us.DB.Find(&UserList)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, UserList)
}

func (us *UserSystem) DeleteUser(c echo.Context) error {
	UserID := c.Param("username")

	var user model.User

	qry := us.DB.Where("username = ?", UserID).First(&user)
	if qry.Error != nil {
		if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data tidak ditemukan",
				"data":    nil,
			})
		}
		c.Echo().Logger.Error("Database error:", qry.Error.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	err := us.DB.Delete(&user).Error
	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada penghapusan data",
			"data":    nil,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (us *UserSystem) UpdateUser(c echo.Context) error {
	userID := c.Param("username")

	var user model.User

	qry := us.DB.Where("username = ?", userID).First(&user)
	if qry.Error != nil {
		if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data tidak ditemukan",
				"data":    nil,
			})
		}
		c.Echo().Logger.Error("Database error:", qry.Error.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	var updateUser model.User

	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	err := us.DB.Model(&user).Updates(&updateUser).Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer updated successfully",
		"data":    &user,
	})
}

func (us *UserSystem) ReadUserById(c echo.Context) error {
	userID := c.Param("username")

	var user model.User

	qry := us.DB.Where("username = ?", userID).First(&user)
	if qry.Error != nil {
		if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data tidak ditemukan",
				"data":    nil,
			})
		}
		c.Echo().Logger.Error("Database error:", qry.Error.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    &user,
	})
}
