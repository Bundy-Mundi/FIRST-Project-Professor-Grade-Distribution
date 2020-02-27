package smc2019spring

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
var fileURL string = "csvfiles/smc/spring_2019.csv"

// RowCount !!
var RowCount = map[string]int{
	"courseName": 2,
	"profName":   3,
	"colA":       4,
	"colB":       5,
	"colC":       6,
	"colD":       7,
	"colF":       8,
	"colNP":      11,
	"colP":       12,
	"colW":       14,
	"colT":       15,
}

// AllData 2019 Spring
func AllData(c echo.Context) error {
	fmt.Println(fileURL)
	spring2019JSON := smcextractor.ExtractJSON(fileURL, RowCount)
	return c.JSONBlob(http.StatusOK, spring2019JSON)
}

// OneData Return One Data
func OneData(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
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

// SearchByProfessor Return Key IDs
func SearchByProfessor(c echo.Context) error {
	result := []int{}
	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	searchTerm := c.Param("name")
	for _, v := range spring2019RAW {
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
	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	searchTerm := c.Param("name")
	for _, v := range spring2019RAW {
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
