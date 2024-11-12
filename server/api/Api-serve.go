package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServeGET(e *echo.Echo) {

	e.GET("/api/v1/parallax/gamer/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

}

func ServePOST(e *echo.Echo) {

	e.POST("/api/v1/parallax/gamer/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

}
