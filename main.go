package main

import (
	"net/http"

	"github.com/Bundy-Mundi/profgradedist/extractor"
	"github.com/labstack/echo"
)

func handleHOME(c echo.Context) error {
	return c.File("home.html")
}

func handle2019Spring(c echo.Context) error {
	spring2019JsonData := extractor.Extract("tabula-spring_2019.csv")
	return c.JSONBlob(http.StatusOK, spring2019JsonData)
}

func main() {

	// Middlewares
	e := echo.New()
	e.GET("/", handleHOME)
	e.GET("/2019/spring", handle2019Spring)

	e.Logger.Fatal(e.Start(":1323"))
}
