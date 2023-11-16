package customer

import (
	"day-22/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	Model model.CustomerQuery
}

func (cc *CustomerController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CreateRequest)

		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(input); err != nil {
			c.Echo().Logger.Error("Input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
		}

		var inputProcess = new(model.CustomerModel)
		inputProcess.Hp = input.Hp
		inputProcess.Nama = input.Nama

		result, err := cc.Model.AddCustomer(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Add Customer, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar pada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(CreateResponse)
		response.Hp = result.Hp
		response.Nama = result.Nama

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (cc *CustomerController) GetListCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := cc.Model.ListCustomer()
		if err != nil {
			c.Logger().Error("ERROR Get Data User, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response []CreateResponse
		for _, customer := range result {
			customerResponse := CreateResponse{
				Nama: customer.Nama,
				Hp:   customer.Hp,
			}
			response = append(response, customerResponse)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get list customer",
			"data":    response,
		})
	}

}

func (cc *CustomerController) GetCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		hp := c.Param("hp")

		result, err := cc.Model.GetCustomerByHp(hp)
		if err != nil {
			c.Logger().Error("ERROR Get Data Customer, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(CreateResponse)
		response.Nama = result.Nama
		response.Hp = result.Hp

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get data customer",
			"data":    response,
		})
	}
}

func (cc *CustomerController) UpdateCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		hp := c.Param("hp")

		var updatedData = new(UpdateRequest)
		if err := c.Bind(updatedData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		validate := validator.New(validator.WithRequiredStructEnabled())

		if err := validate.Struct(updatedData); err != nil {
			c.Echo().Logger.Error("Input error :", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"data":    nil,
			})
		}

		var updatedDataProcess = new(model.CustomerModel)
		updatedDataProcess.Nama = updatedData.Nama

		result, err := cc.Model.UpdateCustomer(hp, *updatedDataProcess)
		if err != nil {
			c.Logger().Error("ERROR Update Data Customer, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(CreateResponse)
		response.Nama = result.Nama
		response.Hp = result.Hp

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success update customer data",
			"data":    response,
		})
	}
}

func (cc *CustomerController) DeleteCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		hp := c.Param("hp")

		err := cc.Model.DeleteCustomer(hp)
		if err != nil {
			c.Logger().Error("ERROR Delete Data Customer, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika menghapus data",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success delete customer data",
		})
	}
}
