package customer

import (
	"day-22/model"
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

func (cs *CustomerSystem) ReadCustomer() ([]model.Customer, bool) {
	var customerList []model.Customer

	qry := cs.DB.Find(&customerList)
	err := qry.Error

	if err != nil {
		fmt.Println("Error read data table:", err.Error())
		return nil, false
	}

	return customerList, true
}

func (cs *CustomerSystem) DeleteCustomer(customerID string) bool {
	var customer model.Customer

	qry := cs.DB.Where("hp = ?", customerID).First(&customer)
	if qry.Error != nil {
		fmt.Println("Customer tidak ditemukan")
		return false
	}

	if err := cs.DB.Delete(&customer).Error; err != nil {
		fmt.Println("Gagal menghapus Customer: ", err.Error())
		return false
	}

	return true
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
