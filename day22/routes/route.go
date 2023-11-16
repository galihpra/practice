package routes

import (
	"day-22/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
}

func routeUser(e *echo.Echo, uc user.UserController) {
	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register())
	e.GET("/users", uc.GetListUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/users/:username", uc.GetUser(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
