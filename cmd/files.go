package cmd

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"github.com/volker48/touchstone/metrics"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Updater interface {
	Update(y, yhat string)
}

type JsonMetrics interface {
	Populate(u *Updater)
}

type ClsJsonMetrics struct {
	*metrics.ConfusionMatrix
	ID        string  `json:"ID"`
	Precision float64 `json:"Precision,omitempty"`
	Recall    float64 `json:"Recall,omitempty"`
	F1        float64 `json:"F1,omitempty"`
	FBeta     float64 `json:"FBeta,omitempty"`
	Beta      float64 `json:"Beta,omitempty"`
	MCC       float64 `json:"MCC,omitempty"`
	YoudenJ   float64 `json:"YoudenJ,omitempty"`
}

type RgrJsonMetrics struct {
	*metrics.Residuals
	ID           string  `json:"ID"`
	MeanY        float64 `json:"MeanY,omitempty"`
	MSE          float64 `json:"MSE,omitempty"`
	RSquared     float64 `json:"RSquared,omitempty"`
}

func (cjm *ClsJsonMetrics) Populate(cm *metrics.ConfusionMatrix) {
	cjm.Precision = cm.Precision()
	cjm.Recall = cm.Recall()
	cjm.F1 = cm.FScore(1.0)
	cjm.FBeta = cm.FScore(beta)
	cjm.Beta = beta
	cjm.MCC = cm.MCC()
	cjm.YoudenJ = cm.YoudenJ()
}

func (rjm *RgrJsonMetrics) Populate(ss *metrics.Residuals) {
	rjm.MeanY = ss.Sum / float64(ss.Count)
	rjm.MSE = ss.MSE()
	rjm.RSquared = ss.RSquared()
}

func dumpJson(jm interface{}, filename string) {
	jsonBytes, err := json.Marshal(jm)
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

func filename2ID(filename string) string {
	base := filepath.Base(filename)
	fileSplit := strings.SplitN(base, ".", 2)
	id := fileSplit[0]
	return id
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
