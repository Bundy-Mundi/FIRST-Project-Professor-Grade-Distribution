package gethome

import (
	"encoding/json"
	"fmt"
	"os"
)

// Data stru
type Data struct {
	Type   string
	Values []map[string]string
	ID     string
}

// GETCollege Fetch Data for Home (currently Fake Data)
func GETCollege() []byte {
	var results []Data
	var values []map[string]string
	v := make(map[string]string)
	v["value"] = "smc"
	v["fullName"] = "Santa Monica College"
	values = append(values, v)
	d := Data{"College", values, "college-select"}
	results = append(results, d)
	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return jsonData
}

// GETSemester s
func GETSemester() []Data {
	// var results []Data
	return nil
}

// GETYear s
func GETYear() []Data {
	// var results []Data
	return nil
}
