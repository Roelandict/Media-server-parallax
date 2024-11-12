package main

import (
	"parallax-website/server/routes"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	routes.Serve(e)

	for {
		time.Sleep(time.Hour)
	}

}
