package routes

import (
	"day-22/controller/customer"
	"day-22/controller/product"
	"day-22/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.UserController, cc customer.CustomerController, pc product.ProductController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeCustomer(e, cc)
	routeProduct(e, pc)
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
	e.GET("/customers", cc.GetListCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/customers/:hp", cc.GetCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.POST("/customers", cc.Create(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/customers/:hp", cc.UpdateCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/customers/:hp", cc.DeleteCustomer(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routeProduct(e *echo.Echo, pc product.ProductController) {
	e.GET("/products", pc.GetListProduct(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/products/:barcode", pc.GetProductByBarcode(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.POST("/products", pc.CreateProduct(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.PUT("/products/:barcode", pc.UpdateProduct(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.DELETE("/products/:barcode", pc.DeleteProduct(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
