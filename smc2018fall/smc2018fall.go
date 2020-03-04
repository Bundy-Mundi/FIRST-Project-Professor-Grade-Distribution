package smc2018fall

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Bundy-Mundi/graderbackend/smcextractor"
	"github.com/labstack/echo"
)

var errNoMatching = errors.New("NO MATCHING ID")

// baseURL
var fileURL string = "csvfiles/smc/fall_2018.csv"

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

	fall2018JSON := smcextractor.ExtractJSON(fileURL, RowCount)
	return c.JSONBlob(http.StatusOK, fall2018JSON)
}

// GetByID Return One Data
func GetByID(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	fall2018RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	for _, v := range fall2018RAW {
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

// SearchByProfessor Return Key IDs
func SearchByProfessor(c echo.Context) error {
	result := []int{}
	fall2018RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	searchTerm := c.Param("name")
	for _, v := range fall2018RAW {
		// Make Terms into Slices
		data := strings.Split(cleanSpace(v.Professor), "")
		search := strings.Split(cleanSpace(searchTerm), "")
		if len(data) > 0 {
			// Check
			ok := isMatch(data, search, 3)
			if ok {
				result = append(result, v.ID)
			}
		}
	}
	// Sending JSON
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, resultJSON)
}

func isMatch(data []string, search []string, restriction int) bool {
	ok := false
	var useShorter int
	matchCounter := 0
	// To Prevent Panic
	if len(search) > len(data) {
		useShorter = len(data)
	} else {
		useShorter = len(search)
	}
	if len(search) > 6 {
		restriction = 6
	}
	for i := 0; i < useShorter; i++ {
		if data[0] == search[0] && data[i] == search[i] {
			matchCounter++
		}
	}
	if matchCounter >= restriction {
		ok = true
	}
	return ok
}

// SearchByClass Return key IDs
func SearchByClass(c echo.Context) error {
	result := []int{}
	fall2018RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	searchTerm := c.Param("name")
	for _, v := range fall2018RAW {
		// Make Terms into Slices
		data := strings.Split(cleanSpace(v.Name), "")
		search := strings.Split(cleanSpace(searchTerm), "")
		if len(data) > 0 {
			// Check
			ok := isMatch(data, search, 2)
			if ok {
				result = append(result, v.ID)
			}
		}
	}
	// Sending JSON
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, resultJSON)
}

func cleanSpace(s string) string {
	return strings.Join(strings.Fields(strings.ToLower(s)), "")
}
