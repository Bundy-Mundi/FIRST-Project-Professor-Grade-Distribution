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

// ProfList 2019 Spring
func ProfList(c echo.Context) error {
	profList := make(map[int]string)
	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
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

// SearchByProfessor Return OneData(id)
func SearchByProfessor(c echo.Context) error {
	matchCounter := 0
	searchTerm := c.Param("name")
	searchTermArray := strings.Split(searchTerm, "")
	profList := make(map[int]string)
	result := make([]int, matchCounter)

	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	for _, v := range spring2019RAW {
		p := v.Professor
		if p == "" {
			p = v.Name
		}
		profList[v.ID] = p
	}
	cleanedList := cleanData(profList)
	for i := 0; i < len(cleanedList); i++ {
		dataArray := strings.Split(cleanedList[i], "")
		if len(dataArray) > 0 {
			for i := 0; i < len(searchTermArray); i++ {
				dataLetter := strings.ToLower(strings.TrimSuffix(dataArray[i], " "))
				searchLetter := strings.ToLower(strings.TrimSuffix(searchTermArray[i], " "))
				if dataLetter == searchLetter {
					matchCounter++
				}
			}
		}
		if matchCounter > 2 {
			result = append(result, i)
		}
		matchCounter = 0
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, resultJSON)
}

// ClassList 2019 Spring
func ClassList(c echo.Context) error {
	classList := make(map[int]string)
	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
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

// SearchByClass Return OneData(id)
func SearchByClass(c echo.Context) error {
	matchCounter := 0
	searchTerm := c.Param("name")
	searchTermArray := strings.Split(searchTerm, "")
	classList := make(map[int]string)
	result := make([]int, matchCounter)

	spring2019RAW := smcextractor.ExtractRAW(fileURL, RowCount)
	for _, v := range spring2019RAW {
		p := v.Professor
		c := v.Name
		if p == "" {
			p = c
		}
		classList[v.ID] = c
	}
	cleanedList := cleanData(classList)
	for i := 0; i < len(cleanedList); i++ {
		dataArray := strings.Split(cleanedList[i], "")
		if len(dataArray) > 0 {
			for i := 0; i < len(searchTermArray); i++ {
				dataLetter := strings.ToLower(strings.TrimSuffix(dataArray[i], " "))
				searchLetter := strings.ToLower(strings.TrimSuffix(searchTermArray[i], " "))
				if dataLetter == searchLetter {
					matchCounter++
				}
			}
		}
		if matchCounter > 2 {
			result = append(result, i)
		}
		matchCounter = 0
	}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c.JSONBlob(http.StatusOK, resultJSON)
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
