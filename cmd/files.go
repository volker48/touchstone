package cmd

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
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
		columnsY := strings.SplitN(yText, " ", 1)
		yHatText := yHatScanner.Text()
		columnsYHat := strings.SplitN(yHatText, " ", 1)
		y, err := strconv.ParseInt(columnsY[0], 10, 8)
		if err != nil {
			log.Fatal("Error parsing int", err)
		}
		yHat, err := strconv.ParseInt(columnsYHat[0], 10, 8)
		if err != nil {
			log.Fatal("Error parsing int", err)
		}
		f(y, yHat)
	}
}
