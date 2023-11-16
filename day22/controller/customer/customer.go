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
