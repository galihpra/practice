package user

import (
	"day-22/model"
	"day-22/utils/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
)

type UserController struct {
	Model model.UserQuery
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterRequest)

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

		var inputProcess = new(model.UserModel)
		inputProcess.Username = input.Username
		inputProcess.Nama = input.Nama
		inputProcess.Password = input.Password
		inputProcess.Status = 2

		result, err := uc.Model.Register(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar pada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(RegisterResponse)
		response.Nama = result.Nama
		response.Username = result.Username

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
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

		result, err := uc.Model.Login(input.Username, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		strToken, err := jwt.GenerateJWT()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.Nama = result.Nama
		response.Username = result.Username
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success login",
			"data":    response,
		})
	}
}

func (uc *UserController) GetListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := uc.Model.ListUser()
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

		var response []UserListResponse
		for _, user := range result {
			userResponse := UserListResponse{
				Nama:     user.Nama,
				Username: user.Username,
				Password: user.Password,
				Status:   user.Status,
			}
			response = append(response, userResponse)
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get list user",
			"data":    response,
		})
	}

}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		result, err := uc.Model.GetUserByUsername(username)
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

		var response = new(UserListResponse)
		response.Nama = result.Nama
		response.Username = result.Username
		response.Password = result.Password
		response.Status = result.Status

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get data user",
			"data":    response,
		})
	}
}
