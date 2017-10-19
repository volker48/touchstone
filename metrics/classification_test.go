// Copyright © 2017 Marcus McCudy <marcus.mccurdy@gmail.com>
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

package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var confusiontests = []struct {
	in  [][]string
	out float64
}{
	{
		in:  [][]string{{"1", "1"}, {"1", "1"}, {"1", "1"}, {"1", "1"}},
		out: 1.0,
	},
	{

		in:  [][]string{{"1", "0"}, {"1", "0"}, {"1", "1"}, {"1", "1"}},
		out: 2.0 / 3.0,
	},
}

func TestConfusionMatrix_F1Score(t *testing.T) {
	for _, ct := range confusiontests {
		cm := &ConfusionMatrix{}
		for i := 0; i < len(ct.in); i++ {
			y := ct.in[i][0]
			yHat := ct.in[i][1]
			cm.Update(y, yHat)

		}
		f1 := cm.FScore(1.0)
		assert.True(t, f1 == ct.out, "Expected %f actual %f precision %f, recall %f TP %f FP %f TN %f FN %f", ct.out, f1, cm.Precision(), cm.Recall(), cm.TP, cm.FP, cm.TN, cm.FN)
	}
}

var thresholdtests = []struct {
	in  [][]string
	out float64
}{
	{
		in:  [][]string{{"0.5", "1"}, {"0.6", "1"}, {"0.7", "1"}, {"0.9", "1"}},
		out: 1.0,
	},
	{

		in:  [][]string{{"0.89", "0"}, {"0.5", "0"}, {"1.0", "1"}, {"0.54", "1"}},
		out: 2.0 / 3.0,
	},
}

func TestConfusionMatrixThreshold_F1Score(t *testing.T) {
	for _, ct := range thresholdtests {
		cm := &ConfusionMatrix{}
		cm.Threshold = 0.5
		for i := 0; i < len(ct.in); i++ {
			y := ct.in[i][0]
			yHat := ct.in[i][1]
			cm.Update(y, yHat)

		}
		f1 := cm.FScore(1.0)
		assert.True(t, f1 == ct.out, "Expected %f actual %f precision %f, recall %f TP %f FP %f TN %f FN %f", ct.out, f1, cm.Precision(), cm.Recall(), cm.TP, cm.FP, cm.TN, cm.FN)
	}

}
