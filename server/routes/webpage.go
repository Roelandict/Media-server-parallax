package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	e.Static("/Fonts", "Web/Fonts")
	e.Static("/css", "Web/css")
	e.Static("/img", "Web/img")

	e.GET("/", func(c echo.Context) error {
		log.Println("Website is served") //debug
		return c.File("Web/pages/index.html")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "testting page")
	})
}
