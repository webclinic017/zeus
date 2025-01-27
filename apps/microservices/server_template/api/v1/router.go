package v1_echo_server_template

import (
	"net/http"
)

func Routes(e *echo.Echo) *echo.Echo {
	// Routes
	e.GET("/health", Health)
	e.GET("/healthy", Health)

	e.GET("/create/entry", EntryRequestHandler)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func EntryRequestHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Sample Response A")
}
