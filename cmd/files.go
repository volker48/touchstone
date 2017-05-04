package cmd

import (
	"os"
	"log"
	"bufio"
	"github.com/volker48/touchstone/metrics"
	"strconv"
)

func readFiles(args []string, f func(int64, int64)) {
	yFile, err := os.Open(args[0])
	if err != nil {
		log.Fatal("You must provide a file for actual values")
	}
	yHatFile, err := os.Open(args[1])
	if err != nil {
		log.Fatal("You must provide a value with predictions")
	}

	yScanner := bufio.NewScanner(bufio.NewReader(yFile))

	yHatScanner := bufio.NewScanner(bufio.NewReader(yHatFile))
	cm := &metrics.ConfusionMatrix{}
	for {
		yScan := yScanner.Scan()
		if !yScan {
			if yScanner.Err() != nil {
				log.Fatal("Scanner error", yScanner.Err())
			}
			break
		}
		yHatScan := yHatScanner.Scan()
		if !yHatScan {
			if yHatScanner.Err() != nil {
				log.Fatal("Scanner error", yHatScanner.Err())
			}
			break

		}

		yText := yScanner.Text()
		yHatText := yHatScanner.Text()
		y, err := strconv.ParseInt(yText, 10, 8)
		if err != nil {
			log.Fatal("Error parsing int", err)
		}
		yHat, err := strconv.ParseInt(yHatText, 10, 8)
		if err != nil {
			log.Fatal("Error parsing int", err)
		}
		f(y, yHat)
		cm.Update(y, yHat)
	}
}
