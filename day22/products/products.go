package products

import (
	"day-22/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductSystem struct {
	DB *gorm.DB
}

func (ps *ProductSystem) CreateProduct(c echo.Context) error {
	var newProduct = new(model.Product)
	if err := c.Bind(newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	qry := ps.DB.Create(newProduct)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error :", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, newProduct)
}

func (ps *ProductSystem) ReadProduct(c echo.Context) error {
	var ProductList []model.Product

	qry := ps.DB.Model(&model.Product{}).
		Select("products.*, users.nama as nama").
		Joins("JOIN users on products.user_id = users.username").
		Scan(&ProductList)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, ProductList)
}

func (ps *ProductSystem) DeleteProduct(c echo.Context) error {
	ProductID := c.Param("barcode")

	var product model.Product

	qry := ps.DB.Where("barcode = ?", ProductID).First(&product)
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

	err := ps.DB.Delete(&product).Error
	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada penghapusan data",
			"data":    nil,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ps *ProductSystem) UpdateProduct(c echo.Context) error {
	ProductID := c.Param("barcode")

	var product model.Product

	qry := ps.DB.Where("barcode = ?", ProductID).First(&product)
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

	var updateProduct model.Product

	if err := c.Bind(&updateProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	err := ps.DB.Model(&product).Updates(&updateProduct).Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer updated successfully",
		"data":    &product,
	})
}

func (ps *ProductSystem) ReadProductById(c echo.Context) error {
	ProductID := c.Param("barcode")

	var product model.Product

	qry := ps.DB.Model(&model.Product{}).
		Where("barcode = ?", ProductID).
		Select("products.*, users.nama as nama").
		Joins("JOIN users on products.user_id = users.username").
		First(&product)
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
		"data":    &product,
	})
}
