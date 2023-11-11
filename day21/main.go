package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required,alpha"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// -------------------- controller --------------------
func CreateUserController(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(User)

		err := c.Bind(input)

		if err != nil {
			c.Echo().Logger.Error("Input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Input data tidak sesuai",
				"data":    nil,
			})
		}

		validate := validator.New(validator.WithRequiredStructEnabled())

		err = validate.Struct(input)

		if err != nil {
			c.Echo().Logger.Error("Input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
		}

		err = conn.Create(input).Error

		if err != nil {
			c.Echo().Logger.Error("Database error :", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "Terjadi kesalahan pada pengolahan data",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "Success",
			"data":    input,
		})
	}
}

func GetUsersController(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []User
		err := conn.Find(&users).Error

		if err != nil {
			c.Echo().Logger.Error("Database error:", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data tidak ditemukan",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success",
			"data":    users,
		})

	}
}

func GetUserController(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")

		var user User
		err := conn.First(&user, userID).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Data tidak ditemukan",
					"data":    nil,
				})
			}

			c.Echo().Logger.Error("Database error:", err.Error())
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
}

func DeleteUserController(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")

		var user User
		err := conn.First(&user, userID).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Data tidak ditemukan",
					"data":    nil,
				})
			}

			c.Echo().Logger.Error("Database error:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi kesalahan pada pengolahan data",
				"data":    nil,
			})
		}

		err = conn.Delete(&user).Error

		if err != nil {
			c.Echo().Logger.Error("Database error:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi kesalahan pada pengolahan data",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User berhasil dihapus",
			"data":    nil,
		})
	}
}

func UpdateUserController(conn *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("id")

		var cek User
		err := conn.First(&cek, userID).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Data user tidak ditemukan",
					"data":    nil,
				})
			}

			c.Echo().Logger.Error("Database error:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
				"data":    nil,
			})
		}

		var updatedUser User
		if err := c.Bind(&updatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"data":    nil,
			})
		}

		err = conn.Model(&cek).Updates(&updatedUser).Error

		if err != nil {
			c.Echo().Logger.Error("Database error:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated successfully",
			"data":    &cek,
		})
	}
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})
	if err != nil {
		e.Logger.Fatal("failed to connect database", err.Error())
	}

	// routing with query parameter
	e.POST("/users", CreateUserController(db))
	e.GET("/users", GetUsersController(db))
	e.GET("/users/:id", GetUserController(db))
	e.DELETE("/users/:id", DeleteUserController(db))
	e.PUT("/users/:id", UpdateUserController(db))

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
