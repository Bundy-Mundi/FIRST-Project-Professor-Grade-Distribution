package gethome

import (
	"encoding/json"
	"fmt"
	"os"
)

// GETCollege Fetch Data for Home (currently Fake Data)
func GETCollege() []byte {
	var values []map[string]string
	v := make(map[string]string)
	v["value"] = "smc"
	v["fullName"] = "Santa Monica College"
	values = append(values, v)
	jsonData, err := json.Marshal(values)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return jsonData
}
