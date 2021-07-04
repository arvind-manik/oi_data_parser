package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"

	. "github.com/artun-o/oi_data_parser/src/models"
)

func parseOIData(filePath string) [][]string {
	file, err := os.Open(filePath)
	fmt.Println(file)
	if err != nil {
		log.Fatal("Unable to read file "+filePath, err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse CSV data "+filePath, err)
	}

	return records
}

const nseDataUrl = "https://www.nseindia.com/api/reports?archives=%5B%7B%22name%22%3A%22F%26O%20-%20Participant%20wise%20Open%20Interest(csv)%22%2C%22type%22%3A%22archives%22%2C%22category%22%3A%22derivatives%22%2C%22section%22%3A%22equity%22%7D%5D%26date%3D{$date}%26type%3Dequity%26mode%3Dsingle"

func main() {
	datasetDir, err := os.Open("./sample_dataset")
	if err != nil {
		log.Fatal("Couldn't read directory ", err)
	}

	defer datasetDir.Close()

	var daywisePositions []MarketwideOI

	fileList, _ := datasetDir.Readdirnames(0)
	for _, name := range fileList {
		records := parseOIData("./sample_dataset/" + name)

		var marketwideOI MarketwideOI
		var participantOI ParticipantOI

		fmt.Printf("%v", reflect.ValueOf(participantOI))

		fmt.Println(records)
	}
}
