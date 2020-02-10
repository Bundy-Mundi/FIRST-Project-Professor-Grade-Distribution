package main

import (
	"net/http"

	"github.com/Bundy-Mundi/profgradedist/extractor"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	jsonData := extractor.Extract()
	s := string(jsonData)
	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, s, "  ")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
