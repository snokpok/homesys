package routes

import "github.com/labstack/echo/v4"

var e *echo.Echo = nil

func Setup() *echo.Echo {
	if e == nil {
		e = echo.New()
	}
	setupOnlineRoutes(e)
	return e
}
