package extractor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

var baseURL = "C://Users/exit2/Desktop/WorkSpace/"

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
func ExtractJSON(fileName string) []byte {
	url := baseURL + fileName
	csvFile, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	results := getAllCors(csvData)
	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return jsonData
}

// ExtractRAW Files
func ExtractRAW(fileName string) []Course {
	url := baseURL + fileName
	csvFile, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	rawData := getAllCors(csvData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rawData
}

func getAllCors(row [][]string) []Course {
	var cors []Course
	var tempCourseName string

	courseName := 2
	profName := 3
	colA := 4
	colB := 5
	colC := 6
	colD := 7
	colF := 8
	colNP := 11
	colP := 12
	colW := 14
	colT := 15

	for i := 1; i < len(row); i++ {
		Name := row[i][courseName]
		if Name != "" {
			tempCourseName = Name
		}
		Professor := row[i][profName]
		Students := row[i][colT]
		Grades := grade{
			A:  row[i][colA],
			B:  row[i][colB],
			C:  row[i][colC],
			D:  row[i][colD],
			F:  row[i][colF],
			NP: row[i][colNP],
			P:  row[i][colP],
			W:  row[i][colW],
		}
		ID := i

		newCourse := Course{tempCourseName, Professor, Grades, Students, ID}
		cors = append(cors, newCourse)
	}
	return cors
}
