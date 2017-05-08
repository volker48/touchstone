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

import "math"

type ConfusionMatrix struct {
	TP    float64
	FP    float64
	TN    float64
	FN    float64
	Total int64
}

func (cm *ConfusionMatrix) Update(y, yHat int64) {
	if y == -1 {
		y = 0
	}
	if yHat == -1 {
		yHat = 0
	}
	switch y {
	case 0:
		switch yHat {
		case 0:
			cm.TN += 1.0
		case 1:
			cm.FP += 1.0
		}
	case 1:
		switch yHat {
		case 0:
			cm.FN += 1.0
		case 1:
			cm.TP += 1.0
		}

	}
	cm.Total++
}

func (cm *ConfusionMatrix) F1Score() float64 {
	p := cm.Precision()
	r := cm.Recall()
	f1 := 2 * (p * r / (p + r))
	return f1
}

func (cm *ConfusionMatrix) Precision() float64 {
	return cm.TP / (cm.TP + cm.FP)
}

func (cm *ConfusionMatrix) Recall() float64 {
	return cm.TP / (cm.TP + cm.FN)
}
func (cm *ConfusionMatrix) MCC() float64 {
	denom := (cm.TP + cm.FP) * (cm.TP + cm.FN) * (cm.TN + cm.FP) * (cm.TN + cm.FN)
	if denom == 0.0 {
		return 0.0
	}
	mcc := ((cm.TP * cm.TN) - (cm.FP * cm.FN)) / math.Sqrt(denom)
	return mcc
}
