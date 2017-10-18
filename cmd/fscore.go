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
	"github.com/spf13/cobra"
	"github.com/volker48/touchstone/metrics"
	"log"
)

var beta float64
var threshold float64

// fscoreCmd represents the fbeta score command
var fscoreCmd = &cobra.Command{
	Use:   "fscore",
	Short: "Calculates F score",
	Long: `Example usage:
	./touchstone fscore y.txt yHat.txt --beta=2.0 --threshold=0.5`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := &metrics.ConfusionMatrix{}
		readFiles(args, cm)
		log.Printf("Total samples: %d", cm.Total)
		log.Printf("Confusion Matrix TP: %d, FP: %d, TN: %d, FN: %d", cm.TP, cm.FP, cm.TN, cm.FN)
		log.Printf("F score beta %f: %f", beta, cm.FScore(beta))
		log.Printf("Matthews correlation coefficient: %f", cm.MCC())
	},
}

func init() {
	RootCmd.AddCommand(fscoreCmd)
	fscoreCmd.Flags().Float64VarP(&beta, "beta", "b", 1.0, "Beta parameter to use when calculating the F score. Defaults to 1.0")
	fscoreCmd.Flags().Float64VarP(&threshold,"threshold", "t", -1.0, "Classification threshold when values in y are probabilities. If set to -1.0 (default), values in y are assumed to be binary.")
}
