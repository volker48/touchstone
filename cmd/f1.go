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
	"github.com/volker48/touchstone/metrics"
	"github.com/spf13/cobra"
	"log"
)



// f1Cmd represents the f1 command
var f1Cmd = &cobra.Command{
	Use:   "f1",
	Short: "Calculates F1 score",
	Long: `Example usage:
	./touchstone f1 y.txt yHat.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		cm := &metrics.ConfusionMatrix{}
		readFiles(args, cm.Update)
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
