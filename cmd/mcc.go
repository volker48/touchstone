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

// mccCmd represents the mcc command
var mccCmd = &cobra.Command{
	Use:   "mcc",
	Short: "Calculates Matthews correlation coefficient",
	Long: `The Matthews correlation coefficient is used in machine learning as a measure of the quality of binary
	(two-class) classifications. It takes into account true and false positives and negatives and is generally
	regarded as a balanced measure which can be used even if the classes are of very different sizes. The MCC is
	in essence a correlation coefficient between the observed and predicted binary classifications; it returns a
	value between −1 and +1.

	A coefficient of:
	+1 represents a perfect prediction,
	 0 no better than random prediction and
	−1 indicates total disagreement between prediction and observation.`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := &metrics.ConfusionMatrix{}
		cm.Threshold = threshold
		readFiles(args, cm)
		log.Printf("Total samples: %d", cm.Total)
		log.Printf("Confusion Matrix TP: %d, FP: %d, TN: %d, FN: %d", cm.TP, cm.FP, cm.TN, cm.FN)
		log.Printf("Matthews correlation coefficient: %f", cm.MCC())
	},
}

func init() {
	ClassificationCmd.AddCommand(mccCmd)
}
