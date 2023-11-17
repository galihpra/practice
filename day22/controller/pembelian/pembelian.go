package pembelian

import (
	"day-22/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PembelianController struct {
	Model model.PembelianQuery
}

func (pbc *PembelianController) Create() echo.HandlerFunc {
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

		var inputProcess = new(model.PembelianModel)
		inputProcess.No_invoice = input.No_invoice
		inputProcess.CustomerID = input.CustomerID
		inputProcess.UserID = input.UserID

		_, err := pbc.Model.AddPembelian(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Add Pembelian, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar pada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
		})
	}
}

func (pbc *PembelianController) GetListPembelian() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := pbc.Model.ListPembelian()
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

		var response []ListResponse
		for _, pembelian := range result {
			pembelianResponse := ListResponse{
				No_invoice: pembelian.No_invoice,
				UserID:     pembelian.UserID,
				CustomerID: pembelian.CustomerID,
				Total:      pembelian.Total,
			}
			response = append(response, pembelianResponse)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get list pembelian",
			"data":    response,
		})
	}

}

func (pbc *PembelianController) GetPembelian() echo.HandlerFunc {
	return func(c echo.Context) error {
		no_invoice := c.Param("no_invoice")

		result, err := pbc.Model.GetPembelianByInvoice(no_invoice)
		if err != nil {
			c.Logger().Error("ERROR Get Data Pembelian, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(ListResponse)
		response.No_invoice = no_invoice
		response.UserID = result.UserID
		response.CustomerID = result.CustomerID
		response.Total = result.Total

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get data pembelian",
			"data":    response,
		})
	}
}

func (pbc *PembelianController) DeletePembelian() echo.HandlerFunc {
	return func(c echo.Context) error {
		no_invoice := c.Param("no_invoice")

		err := pbc.Model.DeletePembelian(no_invoice)
		if err != nil {
			c.Logger().Error("ERROR Delete Data Pembelian, explain:", err.Error())
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
			"message": "success delete pembelian data",
		})
	}
}
