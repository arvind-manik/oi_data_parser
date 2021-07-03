package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func parseOIData(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read file " + filePath, err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	
	if err != nil {
		log.Fatal("Unable to parse CSV data " + filePath, err)
	}

	return records
}

func main() {
	records := parseOIData("./sample_dataset/fao_participant_oi_30062021.csv")
	fmt.Println(records);
}
