package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//instanciar echo
	e := echo.New()

	//ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "holis")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
