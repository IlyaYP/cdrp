package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/IlyaYP/cdrp/model"
)

func main() {
	// open file
	f, err := os.Open("/tmp/cdr.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record, err := model.NewRecordFromStrSlice(rec); err == nil {
			fmt.Println(record.DateTimeOrigination, record.CallingPartyNumber,
				record.OrigDeviceName, record.OriginalCalledPartyNumber, record.DestDeviceName, record.Duration)
		} else {
			log.Println(err)
		}
	}
}
