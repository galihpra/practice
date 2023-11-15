package customer

import (
	"day-22/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerSystem struct {
	DB *gorm.DB
}

func (cs *CustomerSystem) CreateCustomer(c echo.Context) error {
	var newCustomer = new(model.Customer)
	if err := c.Bind(newCustomer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	qry := cs.DB.Create(newCustomer)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error :", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Terjadi kesalahan pada pengolahan data",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, newCustomer)
}

func (cs *CustomerSystem) ReadCustomer(c echo.Context) error {
	var customerList []model.Customer

	qry := cs.DB.Find(&customerList)
	err := qry.Error

	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Data tidak ditemukan",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, customerList)
}

func (cs *CustomerSystem) DeleteCustomer(c echo.Context) error {
	customerID := c.Param("hp")

	var customer model.Customer

	qry := cs.DB.Where("hp = ?", customerID).First(&customer)
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

	err := cs.DB.Delete(&customer).Error
	if err != nil {
		c.Echo().Logger.Error("Database error:", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Terjadi kesalahan pada penghapusan data",
			"data":    nil,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (cs *CustomerSystem) UpdateCustomer(hp string, customerUpdate model.Customer) bool {
	var customer model.Customer
	qry := cs.DB.Where("hp = ?", hp).First(&customer)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	customer.Nama = customerUpdate.Nama

	if err := cs.DB.Model(&customer).Updates(&customer).Error; err != nil {
		fmt.Println("Gagal mengupdate customer: ", err.Error())
		return false
	}

	return true
}
