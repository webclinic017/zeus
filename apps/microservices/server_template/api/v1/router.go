package v1_echo_server_template

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) *echo.Echo {
	// Routes
	e.GET("/health", Health)
	e.GET("/healthy", Health)

	e.GET("/demo", SampleResponse)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func SampleResponse(c echo.Context) error {
	return c.String(http.StatusOK, "Sample Response A")
}
