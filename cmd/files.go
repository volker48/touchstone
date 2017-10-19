package cmd

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Updater interface {
	Update(y, yhat string)
}

type JsonMetrics struct {
	ID        string
	TP        int64
	FP        int64
	TN        int64
	FN        int64
	Total     int64
	Precision float64
	Recall    float64
	F1        float64
	FBeta     float64
	Beta      float64
	MCC       float64
	YoudenJ   float64
}

func dumpJson(metrics JsonMetrics, filename string) {
	jsonBytes, err := json.Marshal(metrics)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}

	jsonFile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Error opening JSON dump file: ", err)
	}
	defer jsonFile.Close()

	if _, err = jsonFile.Write(jsonBytes); err != nil {
		log.Fatal("Error writing JSON to file: ", err)
	}

	log.Printf("JSON data written to: %q", filename)
}

func readFiles(args []string, u Updater) {
	yFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal("You must provide a file for actual values")
	}
	defer yFile.Close()
	yHatFile, err := os.Open(args[1])
	if err != nil {
		log.Fatal("You must provide a value with predictions")
	}
	defer yHatFile.Close()
	yExt := filepath.Ext(args[0])
	yHatExt := filepath.Ext(args[1])
	var yScanner *bufio.Scanner
	var yHatScanner *bufio.Scanner

	if yExt == ".gz" {
		gzr, err := gzip.NewReader(yFile)
		if err != nil {
			log.Fatal("Couldn't reader gzip file: ", err)
		}
		defer gzr.Close()
		yScanner = bufio.NewScanner(bufio.NewReader(gzr))
	} else {
		yScanner = bufio.NewScanner(bufio.NewReader(yFile))
	}

	if yHatExt == ".gz" {
		gzr, err := gzip.NewReader(yHatFile)
		if err != nil {
			log.Fatal("Couldn't read gzip file: ", err)
		}
		defer gzr.Close()
		yHatScanner = bufio.NewScanner(bufio.NewReader(gzr))
	} else {
		yHatScanner = bufio.NewScanner(bufio.NewReader(yHatFile))
	}

	for {
		yScan := yScanner.Scan()
		yHatScan := yHatScanner.Scan()

		if !yScan || !yHatScan {
			if yScanner.Err() != nil {
				log.Fatal("Scanner error", yScanner.Err())
			}
			if yHatScanner.Err() != nil {
				log.Fatal("Scanner error", yHatScanner.Err())
			}

			if !yScan && yHatScan {

				log.Fatal("File of actual labels is shorter than file of predictions.")
			}

			if yScan && !yHatScan {

				log.Fatal("File of predictions is shorter than file of labels.")
			}

			break
		}

		yText := yScanner.Text()
		columnsY := strings.SplitN(yText, " ", 2)
		yHatText := yHatScanner.Text()
		columnsYHat := strings.SplitN(yHatText, " ", 2)

		u.Update(columnsY[0], columnsYHat[0])
	}
}
