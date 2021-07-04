package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func parseOIData(filePath string) [][]string {
	file, err := os.Open(filePath)
	fmt.Println(file)
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
	datasetDir, err := os.Open("./sample_dataset")
	if err != nil {
		log.Fatal("Couldn't read directory ", err)
	}

	defer datasetDir.Close()

	fileList, _ := datasetDir.Readdirnames(0)
	for _, name := range fileList {
		fmt.Println("File name: %s, Kind: %s", name, reflect.TypeOf(name));
	}

	records := parseOIData("./sample_dataset/fao_participant_oi_30062021.csv")
	fmt.Println(records);
}
