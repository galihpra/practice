package pembelian

import (
	"day-22/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PembelianSystem struct {
	DB *gorm.DB
}

func (ps *PembelianSystem) CreatePembelian(c echo.Context) error {
	var newPembelian = new(model.Pembelian)
	if err := c.Bind(newPembelian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	qry := ps.DB.Create(newPembelian)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error :", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, newPembelian)
}

func (ps *PembelianSystem) ReadPembelian(c echo.Context) error {
	var PembelianList []model.Pembelian

	qry := ps.DB.Find(&PembelianList)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, PembelianList)
}

func (ps *PembelianSystem) DeletePembelian(c echo.Context) error {
	PembelianID := c.Param("no_invoice")

	var pembelian model.Pembelian

	qry := ps.DB.Where("no_invoice = ?", PembelianID).First(&pembelian)
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

	err := ps.DB.Delete(&pembelian).Error
	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada penghapusan data",
			"data":    nil,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ps *PembelianSystem) UpdatePembelian(c echo.Context) error {
	PembelianID := c.Param("no_invoice")

	var pembelian model.Pembelian

	qry := ps.DB.Where("no_invoice = ?", PembelianID).First(&pembelian)
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

	var updatePembelian model.Pembelian

	if err := c.Bind(&updatePembelian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	err := ps.DB.Model(&pembelian).Updates(&updatePembelian).Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer updated successfully",
		"data":    &pembelian,
	})
}

func (ps *PembelianSystem) ReadPembelianById(c echo.Context) error {
	PembelianID := c.Param("no_invoice")

	var pembelian model.Pembelian

	qry := ps.DB.Where("no_invoice = ?", PembelianID).First(&pembelian)
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
		"data":    &pembelian,
	})
}
