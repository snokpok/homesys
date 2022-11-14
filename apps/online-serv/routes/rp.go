package routes

func setupRPRoutes(e *echo.Echo) {
	e.GET("/rp/health")
}
