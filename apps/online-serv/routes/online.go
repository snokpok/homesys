package routes

import (
	"homesys/apps/online-serv/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupOnlineRoutes(e *echo.Echo) {
	e.GET("/online/health", func(c echo.Context) error {
		rsp := responses.SuccessfulResponse{
			Meta: responses.ResponseMeta{
				Msg: "Message response successful",
			},
			Data: echo.Map{
				"hello": "world",
			},
		}
		return c.JSON(http.StatusAccepted, rsp)
	})
}
