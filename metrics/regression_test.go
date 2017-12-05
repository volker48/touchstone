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

var residualsTestData = []struct {
	in    [][]string
	inLog [][]string
	mse   float64
	r2    float64
}{
	{
		in:    [][]string{{"3", "2.5"}, {"0.5", "1"}, {"2", "2"}, {"7", "8"}},
		inLog: [][]string{{"1.09861229", "0.91629073"}, {"-0.69314718", "0."}, {"0.69314718", "0.69314718"}, {"1.94591015", "2.07944154"}},
		mse:   0.375,
		r2:    0.93530997304582209,
	},
	{
		in:    [][]string{{"0.001", "0.0015"}, {"1.1", "0.99"}, {"0.01", "0.01"}, {"0.000005", "0.000008"}},
		inLog: [][]string{{"-6.90775528", "-6.50229017"}, {"0.09531018", "-0.01005034"}, {"-4.60517019", "-4.60517019"}, {"-12.20607265", "-11.73606902"}},
		mse:   0.0030250625022500057,
		r2:    0.98657791593971977,
	},
}

func TestResiduals_MSE(t *testing.T) {

	for _, rmt := range residualsTestData {
		ss := &Residuals{}
		ss.LogTransform = false
		for i := 0; i < len(rmt.in); i++ {
			y := rmt.in[i][0]
			yHat := rmt.in[i][1]
			ss.Update(y, yHat)
		}
		mse := ss.MSE()
		assert.InDelta(t, rmt.mse, mse, 0.000001, "Expected %f\nActual  %f", rmt.mse, mse)
	}
}

func TestResiduals_MSE_w_LogTransform(t *testing.T) {

	for _, rmt := range residualsTestData {
		ss := &Residuals{}
		ss.LogTransform = true
		for i := 0; i < len(rmt.inLog); i++ {
			y := rmt.inLog[i][0]
			yHat := rmt.inLog[i][1]
			ss.Update(y, yHat)
		}
		mse := ss.MSE()
		assert.InDelta(t, rmt.mse, mse, 0.000001, "Expected %f\nActual  %f", rmt.mse, mse)
	}
}

func TestResiduals_RSquared(t *testing.T) {
	for _, rrt := range residualsTestData {
		ss := &Residuals{}
		ss.LogTransform = false
		for i := 0; i < len(rrt.in); i++ {
			y := rrt.in[i][0]
			yHat := rrt.in[i][1]
			ss.Update(y, yHat)
		}
		r2 := ss.RSquared()
		assert.InDelta(t, rrt.r2, r2, 0.000001, "Expected %f\nActual  %f", rrt.r2, r2)
	}
}

func TestResiduals_RSquared_w_LogTransform(t *testing.T) {
	for _, rrt := range residualsTestData {
		ss := &Residuals{}
		ss.LogTransform = true
		for i := 0; i < len(rrt.inLog); i++ {
			y := rrt.inLog[i][0]
			yHat := rrt.inLog[i][1]
			ss.Update(y, yHat)
		}
		r2 := ss.RSquared()
		assert.InDelta(t, rrt.r2, r2, 0.000001, "Expected %f\nActual  %f", rrt.r2, r2)
	}
}
