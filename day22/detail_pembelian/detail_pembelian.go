package detailpembelian

import (
	"day-22/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DetailPembelianSystem struct {
	DB *gorm.DB
}

func (dps *DetailPembelianSystem) CreateDetailPembelian(c echo.Context) error {
	var newDetailPembelian = new(model.DetailPembelian)
	if err := c.Bind(newDetailPembelian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	qry := dps.DB.Create(newDetailPembelian)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error :", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, newDetailPembelian)
}

func (dps *DetailPembelianSystem) ReadDetailPembelian(c echo.Context) error {
	var detailPembelianList []model.DetailPembelian

	qry := dps.DB.Find(&detailPembelianList)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, detailPembelianList)
}

func (dps *DetailPembelianSystem) DeleteDetailPembelian(c echo.Context) error {
	DetailPembelianID := c.Param("pembelian_id")

	var detailPembelian model.DetailPembelian

	qry := dps.DB.Where("pembelian_id = ?", DetailPembelianID).First(&detailPembelian)
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

	err := dps.DB.Delete(&detailPembelian).Error
	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada penghapusan data",
			"data":    nil,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (dps *DetailPembelianSystem) UpdateDetailPembelian(c echo.Context) error {
	DetailPembelianID := c.Param("pembelian_id")

	var detailPembelian model.DetailPembelian

	qry := dps.DB.Where("pembelian_id = ?", DetailPembelianID).First(&detailPembelian)
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

	var updateDetailPembelian model.Pembelian

	if err := c.Bind(&updateDetailPembelian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"data":    nil,
		})
	}

	err := dps.DB.Model(&detailPembelian).Updates(&updateDetailPembelian).Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Customer updated successfully",
		"data":    &detailPembelian,
	})
}

func (dps *DetailPembelianSystem) ReadDetailPembelianById(c echo.Context) error {
	DetailPembelianID := c.Param("pembelian_id")

	var detailPembelian model.DetailPembelian

	qry := dps.DB.Where("pembelian_id = ?", DetailPembelianID).First(&detailPembelian)
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
		"data":    &detailPembelian,
	})
}
