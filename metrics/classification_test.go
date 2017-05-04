package metrics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var confusiontests = []struct {
	in  [][]int64
	out float64
}{
	{
		in: [][]int64{[]int64{1, 1}, []int64{1, 1}, []int64{1, 1}, []int64{1, 1}},
		out: 1.0,
	},
	{

		in: [][]int64{[]int64{1, 0}, []int64{1, 0}, []int64{1, 1}, []int64{1, 1}},
		out: 2.0/3.0,
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
		f1 := cm.F1Score()
		assert.True(t, f1 == ct.out, "Expected %f actual %f precision %f, recall %f TP %f FP %f TN %f FN %f", ct.out, f1, cm.Precision(), cm.Recall(), cm.TP, cm.FP, cm.TN, cm.FN)
	}

}
