package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type ParticipantOI struct {
	future_index_long int
	future_index_short int

	future_stock_long int
	future_stock_short int

	option_index_call_long int
	option_index_put_long int

	option_index_call_short int
	option_index_put_short int

	option_stock_call_long int
	option_stock_put_long int

	option_stock_call_short int
	option_stock_put_short int
}

type MarketwideOI struct {
	position_date time.Time

	client ParticipantOI
	dii ParticipantOI
	fii ParticipantOI
	pro ParticipantOI
}

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
		records := parseOIData("./sample_dataset/" + name)



		fmt.Println(records);
	}
}
