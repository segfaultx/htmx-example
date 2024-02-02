package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type TableData struct {
	Id      int
	Company string `form:"company"`
	Contact string `form:"contact"`
	Country string `form:"country"`
}

var Tabledata = getData()

func getData() []TableData {
	file, err := os.Open("data.json")

	if err != nil {
		log.Fatalln("Error opening data")
	}
	bytes, err := io.ReadAll(file)

	if err != nil {
		log.Fatalln("Error reading data bytes")
	}

	var data = []TableData{}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		log.Fatalln("Error unmarshalling data")
	}

	return data
}
