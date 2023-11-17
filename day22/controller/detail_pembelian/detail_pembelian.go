package detailpembelian

import (
	"day-22/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DetailPembelianController struct {
	Model model.DetailPembelianQuery
}

func (dpc *DetailPembelianController) Create() echo.HandlerFunc {
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

		var inputProcess = new(model.DetailPembelianModel)
		inputProcess.PembelianID = input.PembelianID
		inputProcess.ProductID = input.ProductID
		inputProcess.Qty = input.Qty

		result, err := dpc.Model.AddDetailPembelian(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Add Detail Pembelian, explain:", err.Error())
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
		response.PembelianID = result.PembelianID
		response.Qty = result.Qty
		response.ProductID = result.ProductID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sudpcess create data",
			"data":    response,
		})
	}
}

func (dpc *DetailPembelianController) GetListDetailPembelian() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := dpc.Model.ListDetailPembelian()
		if err != nil {
			c.Logger().Error("ERROR Get Data Detail Pembelian, explain:", err.Error())
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
		for _, DetailPembelian := range result {
			DetailPembelianResponse := CreateResponse{
				PembelianID: DetailPembelian.PembelianID,
				ProductID:   DetailPembelian.ProductID,
				Qty:         DetailPembelian.Qty,
			}
			response = append(response, DetailPembelianResponse)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "sudpcess get list Detail Pembelian",
			"data":    response,
		})
	}

}

func (dpc *DetailPembelianController) DeleteDetailPembelian() echo.HandlerFunc {
	return func(c echo.Context) error {
		pembelian_id := c.Param("pembelian_id")

		err := dpc.Model.DeleteDetailPembelian(pembelian_id)
		if err != nil {
			c.Logger().Error("ERROR Delete Data Detail Pembelian, explain:", err.Error())
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
			"message": "sudpcess delete Detail Pembelian data",
		})
	}
}
