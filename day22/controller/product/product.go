package product

import (
	"day-22/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Model model.ProductQuery
}

func (pc *ProductController) CreateProduct() echo.HandlerFunc {
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

		var inputProcess = new(model.ProductModel)
		inputProcess.Barcode = input.Barcode
		inputProcess.Nama = input.Nama
		inputProcess.Stok = input.Stok
		inputProcess.Harga = input.Harga
		inputProcess.UserID = input.UserID

		result, err := pc.Model.AddProduct(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Add Product, explain:", err.Error())
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
		response.Barcode = result.Barcode
		response.Nama = result.Nama
		response.Stok = result.Stok
		response.Harga = result.Harga
		response.UserID = result.UserID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (pc *ProductController) GetListProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := pc.Model.ListProduct()
		if err != nil {
			c.Logger().Error("ERROR Get Data Product, explain:", err.Error())
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
		for _, product := range result {
			productResponse := ListResponse{
				Barcode:  product.Barcode,
				Nama:     product.Nama,
				Stok:     product.Stok,
				Harga:    product.Harga,
				UserNama: product.UserNama,
			}
			response = append(response, productResponse)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get list product",
			"data":    response,
		})
	}

}

func (pc *ProductController) GetProductByBarcode() echo.HandlerFunc {
	return func(c echo.Context) error {
		barcode := c.Param("barcode")

		result, err := pc.Model.GetProductByBarcode(barcode)
		if err != nil {
			c.Logger().Error("ERROR Get Data Product, explain:", err.Error())
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
		response.Barcode = barcode
		response.Nama = result.Nama
		response.Stok = result.Stok
		response.Harga = result.Harga
		response.UserNama = result.UserNama

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get data product",
			"data":    response,
		})
	}
}

func (pc *ProductController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		barcode := c.Param("barcode")

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

		var updatedDataProcess = new(model.ProductModel)
		updatedDataProcess.Nama = updatedData.Nama
		updatedDataProcess.Stok = updatedData.Stok
		updatedDataProcess.Harga = updatedData.Harga
		updatedDataProcess.UserID = updatedData.UserID

		result, err := pc.Model.UpdateProduct(barcode, *updatedDataProcess)
		if err != nil {
			c.Logger().Error("ERROR Update Data Product, explain:", err.Error())
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
		response.Barcode = barcode
		response.Nama = result.Nama
		response.Stok = result.Stok
		response.Harga = result.Harga
		response.UserNama = result.UserNama

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success update product data",
			"data":    response,
		})
	}
}

func (pc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		barcode := c.Param("barcode")

		err := pc.Model.DeleteProduct(barcode)
		if err != nil {
			c.Logger().Error("ERROR Delete Data Product, explain:", err.Error())
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
			"message": "success delete product data",
		})
	}
}
