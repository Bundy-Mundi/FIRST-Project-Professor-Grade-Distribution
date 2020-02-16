package main

import (
	"net/http"

	"github.com/Bundy-Mundi/profgradedist/gethome"
	"github.com/Bundy-Mundi/profgradedist/smc2019spring"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HOME
func getCollege(c echo.Context) error {
	data := gethome.GETCollege()
	return c.JSONBlob(http.StatusOK, data)
}
func getYear(c echo.Context) error {
	return nil
}
func getSemester(c echo.Context) error {

	return nil
}

// 2019
func get2019(c echo.Context) error {
	return c.File("templates/2019S.html")
}

// 2019 Fall
func get2019Fall(c echo.Context) error {
	return c.File("templates/2019F.html")
}
func get2019FallProfList(c echo.Context) error {
	return nil
}
func get2019FallClassList(c echo.Context) error {
	return nil
}

// GET 2018

func get2018Spring(c echo.Context) error {
	return c.File("templates/2018S.html")
}
func get2018Fall(c echo.Context) error {
	return c.File("templates/2018F.html")
}

// POST 2019
func smcPOST2019Spring(c echo.Context) error {
	// yaer := c.QueryParam("year")
	// semester := c.QueryParam("semester")
	// searchBy := c.QueryParam("search-by")

	return nil
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CORS())

	// Common API
	e.GET("/api/v1/college", getCollege)
	e.GET("/api/v1/year", getYear)
	e.GET("/api/v1/semester", getSemester)

	// SMC GET
	e.GET("/api/v1/smc/2019", get2019)
	e.GET("/api/v1/smc/2019/spring", smc2019spring.AllData)
	e.GET("/api/v1/smc/2019/spring/:id", smc2019spring.OneData)
	e.GET("/api/v1/smc/2019/spring/prof", smc2019spring.ProfList)
	e.GET("/api/v1/smc/2019/spring/class", smc2019spring.ClassList)
	e.GET("/api/v1/smc/2019/fall", get2019Fall)
	e.GET("/api/v1/smc/2019/fall/:id", get2019Fall)
	e.GET("/api/v1/smc/2019/fall/prof", get2019FallProfList)
	e.GET("/api/v1/smc/2019/fall/class", get2019FallClassList)

	e.Logger.Fatal(e.Start(":1323"))
}
