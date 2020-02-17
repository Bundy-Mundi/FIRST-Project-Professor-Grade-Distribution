package main

import (
	"net/http"

	"github.com/Bundy-Mundi/profgradedist/gethome"
	"github.com/Bundy-Mundi/profgradedist/smc2018fall"
	"github.com/Bundy-Mundi/profgradedist/smc2019spring"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HOME
func getCollege(c echo.Context) error {
	data := gethome.GETCollege()
	return c.JSONBlob(http.StatusOK, data)
}

// SMC 2019
func smc2019(c echo.Context) error {
	return c.File("templates/2019S.html")
}

// SMC 2019 Fall
func get2019Fall(c echo.Context) error {
	return c.File("templates/2019F.html")
}
func get2019FallProfList(c echo.Context) error {
	return nil
}
func get2019FallClassList(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CORS())

	// Common API
	e.GET("/api/v1/college", getCollege)

	// SMC

	e.GET("/api/v1/smc/2019", smc2019)
	//e.GET("/api/v1/smc/2018", get2018)

	// SMC 2019 Spring
	e.GET("/api/v1/smc/2019/spring", smc2019spring.AllData)
	e.GET("/api/v1/smc/2019/spring/:id", smc2019spring.OneData)
	e.GET("/api/v1/smc/2019/spring/prof", smc2019spring.ProfList)
	e.GET("/api/v1/smc/2019/spring/class", smc2019spring.ClassList)

	// SMC 2019 Fall (yet)
	// e.GET("/api/v1/smc/2019/fall", get2019Fall)
	// e.GET("/api/v1/smc/2019/fall/:id", get2019Fall)
	// e.GET("/api/v1/smc/2019/fall/prof", get2019FallProfList)
	// e.GET("/api/v1/smc/2019/fall/class", get2019FallClassList)

	// SMC 2018 Spring (csv file is not prepared)
	//e.GET("/api/v1/smc/2018/spring", smc2018spring.AllData)
	//e.GET("/api/v1/smc/2018/spring/:id", smc2018spring.OneData)
	//e.GET("/api/v1/smc/2018/spring/prof", smc2018spring.ProfList)
	//e.GET("/api/v1/smc/2018/spring/class", smc2018spring.ClassList)

	// SMC 2018 Fall
	e.GET("/api/v1/smc/2018/fall", smc2018fall.AllData)
	e.GET("/api/v1/smc/2018/fall/:id", smc2018fall.OneData)
	e.GET("/api/v1/smc/2018/fall/prof", smc2018fall.ProfList)
	e.GET("/api/v1/smc/2018/fall/class", smc2018fall.ClassList)

	e.Logger.Fatal(e.Start(":1323"))
}
