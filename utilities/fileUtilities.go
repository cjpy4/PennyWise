package utilities

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// var mySlice []int = []int{1,2,3,4,5,6}

func getFile(filePath string) *os.File {
	fmt.Println("filePath:",filePath)
	file, err := os.Open(filePath)
	if (err != nil) {
		fmt.Println("Error opening file:",err)
		return nil
	}
	return file
}


func csvToJSON(csvFile os.File) []map[string]any {
	csvReader := csv.NewReader(&csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("cannot retrieve records from Reader", err.Error())
	}
	var headers []string = []string{}
	var jsonArray []map[string]any
	for i, r := range records {
		if (i == 0) {
			headers = r
		} else {
			rowMap := make(map[string]any)
			for j, v := range r {
				rowMap[headers[j]] = v
			}
			jsonArray = append(jsonArray, rowMap)
		}
	}

	return jsonArray
}

func stringifyJSON(jsonArray []map[string]any) json.RawMessage {
	output, err := json.Marshal(jsonArray)
	if (err != nil) {
		fmt.Println("Cannot encode jsonArray as JSON", err)
		return nil
	}

	return output
}