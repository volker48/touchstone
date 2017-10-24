package metrics

import (
	"log"
	"math"
	"strconv"
)

type Residuals struct {
	Sum_Res      float64   `json:"SumResiduals"`
	SS_Res       float64   `json:"SquaredSumResiduals"`
	Sum          float64   `json:"SumY"`
	Count        int64     `json:"Total"`
	LogTransform bool      `json:"LogTransform"`
	ys           []float64 `json:"-"`
	yHats        []float64 `json:"-"`
}

func (ss *Residuals) RSquared() float64 {
	average := ss.Sum / float64(ss.Count)
	var ssTot float64
	for _, y := range ss.ys {
		tot := y - average
		ssTot += tot * tot
	}
	return 1.0 - (ss.SS_Res / ssTot)
}

func (ss *Residuals) MSE() float64 {
	return ss.SS_Res / float64(ss.Count)
}

func (ss *Residuals) ExplainedVariance() float64 {
	y_diff_avg := ss.Sum_Res / float64(ss.Count)
	y_true_avg := ss.Sum / float64(ss.Count)
	var output float64
	for i, y := range ss.ys {
		n_term := y - ss.yHats[i] - y_diff_avg
		numerator := n_term * n_term
		d_term := y - y_true_avg
		denominator := d_term * d_term
		output += 1 - (numerator / denominator)
	}
	return output / float64(ss.Count)
}

func (ss *Residuals) Update(yText, yHatText string) {
	y, err := strconv.ParseFloat(yText, 64)
	if err != nil {
		log.Fatal("Error parsing y value as float: ", err)
	}

	yHat, err := strconv.ParseFloat(yHatText, 64)
	if err != nil {
		log.Fatal("Error parsing yhat value as float: ", err)
	}

	if ss.LogTransform {
		y = math.Exp(y)
		yHat = math.Exp(yHat)
	}

	ss.Count++
	ss.Sum += y
	e := y - yHat
	ss.Sum_Res += e
	ss.SS_Res += e * e
	ss.ys = append(ss.ys, y)
	ss.yHats = append(ss.yHats, yHat)
}
