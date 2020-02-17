package smc2018fall

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Bundy-Mundi/profgradedist/smcextractor"
	"github.com/labstack/echo"
)

var errNoMatching = errors.New("NO MATCHING ID")

// baseURL
var baseURL string = "tabula-fall_2018.csv"

// RowCount !!
var RowCount = map[string]int{
	"courseName": 0,
	"profName":   1,
	"colA":       2,
	"colB":       3,
	"colC":       4,
	"colD":       5,
	"colF":       6,
	"colNP":      9,
	"colP":       10,
	"colW":       12,
	"colT":       13,
}

// AllData 2018 FallFall
func AllData(c echo.Context) error {

	spring2019JSON := smcextractor.ExtractJSON(baseURL, RowCount)
	return c.JSONBlob(http.StatusOK, spring2019JSON)
}

// OneData Return One Data
func OneData(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	spring2019RAW := smcextractor.ExtractRAW(baseURL, RowCount)
	for _, v := range spring2019RAW {
		if v.ID == ID {
			oneJSON, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return c.JSONBlob(http.StatusOK, oneJSON)
		}
	}
	return errNoMatching
}

// ProfList 2018 Fall
func ProfList(c echo.Context) error {
	profList := make(map[int]string)
	spring2019RAW := smcextractor.ExtractRAW(baseURL, RowCount)
	for _, v := range spring2019RAW {
		p := v.Professor
		if p == "" {
			p = v.Name
		}
		profList[v.ID] = p
	}
	cleanedList := cleanData(profList)
	profJSON, err := json.Marshal(cleanedList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, profJSON)
}

// ClassList 2018 Fall
func ClassList(c echo.Context) error {
	classList := make(map[int]string)
	spring2019RAW := smcextractor.ExtractRAW(baseURL, RowCount)
	for _, v := range spring2019RAW {
		c := v.Name
		classList[v.ID] = c
	}
	cleanedList := cleanData(classList)
	classJSON, err := json.Marshal(cleanedList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, classJSON)
}

func cleanData(data map[int]string) map[int]string {
	for i := 0; i < len(data); i++ {
		d := data[i]
		for i2 := 0; i2 < len(data); i2++ {
			check := data[i2]
			if i != i2 && d == check {
				delete(data, i2)
			}
		}
	}
	return data
}
