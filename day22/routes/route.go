package routes

import (
	"day-22/controller/customer"
	"day-22/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.UserController, cc customer.CustomerController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeCustomer(e, cc)
}

func routeUser(e *echo.Echo, uc user.UserController) {
	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register())
	e.GET("/users", uc.GetListUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/users/:username", uc.GetUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/users/:username", uc.UpdateUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/users/:username", uc.DeleteUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routeCustomer(e *echo.Echo, cc customer.CustomerController) {
	// e.GET("/customers", cc.GetListCustomers(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.GET("/customers/:id", cc.GetCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.POST("/customers", cc.Create(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.PUT("/customers/:id", cc.UpdateCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.DELETE("/customers/:id", cc.DeleteCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
