// Copyright © 2017 Marcus McCurdy <marcus.mccurdy@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/volker48/touchstone/metrics"
	"github.com/spf13/cobra"
	"os"
	"log"
	"bufio"
	"strconv"
)



// f1Cmd represents the f1 command
var f1Cmd = &cobra.Command{
	Use:   "f1",
	Short: "Calculates F1 score",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("f1 called")
		yFile, err := os.Open(args[0])
		if err != nil {
			log.Fatal("You must provide a file for actual values")
		}
		yHatFile, err := os.Open(args[1])
		if err != nil {
			log.Fatal("You must provide a value with preidctions")
		}
		defer yFile.Close()
		defer yHatFile.Close()

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
			cm.Update(y, yHat)
		}
		log.Printf("Total samples: %d", cm.Total)
		log.Printf("Confusion Matrix TP: %f, FP: %f, TN: %f, FN: %f", cm.TP, cm.FP, cm.TN, cm.FN)
		log.Printf("F1 score: %f", cm.F1Score())
	},
}

func init() {
	RootCmd.AddCommand(f1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// f1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// f1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
