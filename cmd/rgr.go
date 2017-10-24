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
	"path/filepath"
	"strings"
)

var log_transform bool

var RegressionCmd = &cobra.Command{
	Use:   "rgr",
	Short: "Calculates a full slate of regression metrics",
	Long: `Example usage:
	./touchstone rgr y.txt yHat.txt -j=ts.json`,
	Run: func(cmd *cobra.Command, args []string) {
		residuals := &metrics.Residuals{}
		residuals.LogTransform = log_transform
		readFiles(args, residuals)

		rjm := &RgrJsonMetrics{}
		rjm.Residuals = residuals
		rjm.Populate(residuals)

		if jsonFile != "" {
			base := filepath.Base(args[1])
			fileSplit := strings.SplitN(base, ".", 2)
			rjm.ID = fileSplit[0]
			dumpJson(rjm, jsonFile)
		}

		log.Printf("Number of samples: %d", residuals.Count)
		log.Printf("Mean y: %f", rjm.MeanY)
		log.Printf("Mean Squared Error: %f", rjm.MSE)
		log.Printf("R Squared: %f", rjm.RSquared)
		log.Printf("Explained Variance: %f", rjm.ExplainedVar)
	},
}

func init() {
	RootCmd.AddCommand(RegressionCmd)
	RegressionCmd.PersistentFlags().BoolVarP(&log_transform, "log_transform", "l", false, "Set this flag if true and predicted values are log transformed in order to inverse the transform for metrics calculations.")
}
