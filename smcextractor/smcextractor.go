package smcextractor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Course Structure
type Course struct {
	Name      string
	Professor string
	Grades    grade
	Students  string
	ID        int
}

type grade struct {
	A  string
	B  string
	C  string
	D  string
	F  string
	NP string
	P  string
	W  string
}

// ExtractJSON Files
func ExtractJSON(fileName string, rowCount map[string]int) []byte {

	// Get the CSV File's Directory
	url, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// Open the CSV File
	csvFile, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	// Store Data Into []Map
	results := getAllCors(csvData, rowCount)

	// Convert Data Into a Json
	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return jsonData
}

// ExtractRAW Files
func ExtractRAW(fileName string, rowCount map[string]int) []Course {
	url, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
	}
	csvFile, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	rawData := getAllCors(csvData, rowCount)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rawData
}

func getAllCors(row [][]string, rowCount map[string]int) []Course {
	var cors []Course
	var tempCourseName string
	for i := 1; i < len(row); i++ {
		Name := row[i][rowCount["courseName"]]
		if Name != "" {
			tempCourseName = Name
		}
		Professor := row[i][rowCount["profName"]]
		Students := row[i][rowCount["colT"]]
		Grades := grade{
			A:  row[i][rowCount["colA"]],
			B:  row[i][rowCount["colB"]],
			C:  row[i][rowCount["colC"]],
			D:  row[i][rowCount["colD"]],
			F:  row[i][rowCount["colF"]],
			NP: row[i][rowCount["colNP"]],
			P:  row[i][rowCount["colP"]],
			W:  row[i][rowCount["colW"]],
		}
		ID := i
		newCourse := Course{tempCourseName, Professor, Grades, Students, ID}
		cors = append(cors, newCourse)
		// fmt.Println(Professor, Students, Grades)
	}
	return cors
}
