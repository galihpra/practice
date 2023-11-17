package routes

import (
	"day-22/controller/customer"
	detailpembelian "day-22/controller/detail_pembelian"
	"day-22/controller/pembelian"
	"day-22/controller/product"
	"day-22/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.UserController, cc customer.CustomerController,
	pc product.ProductController, pbc pembelian.PembelianController, dpc detailpembelian.DetailPembelianController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
	routeCustomer(e, cc)
	routeProduct(e, pc)
	routePembelian(e, pbc)
	routeDetailPembelian(e, dpc)
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

func routePembelian(e *echo.Echo, pbc pembelian.PembelianController) {
	e.GET("/pembelians", pbc.GetListPembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.GET("/pembelians/:no_invoice", pbc.GetPembelianByInvoice(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.POST("/pembelians", pbc.Create(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.PUT("/pembelians/:barcode", pbc.UpdatePembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.DELETE("/pembelians/:barcode", pbc.DeletePembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}

func routeDetailPembelian(e *echo.Echo, dpc detailpembelian.DetailPembelianController) {
	// e.GET("/details", dpc.GetListDetrouteDetailPembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.GET("/details/:barcode", dpc.GetDetrouteDetailPembelianByBarcode(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.POST("/details", dpc.Create(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.PUT("/details/:barcode", dpc.UpdateDetrouteDetailPembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	// e.DELETE("/details/:barcode", dpc.DeleteDetrouteDetailPembelian(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
