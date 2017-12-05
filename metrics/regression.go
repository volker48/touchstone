package metrics

import (
	"log"
	"math"
	"strconv"
)

type Residuals struct {
	SSRes        float64   `json:"SquaredSumResiduals"`
	Sum          float64   `json:"SumY"`
	Count        int64     `json:"Total"`
	LogTransform bool      `json:"LogTransform"`
	ys           []float64 `json:"-"`
}

func (ss *Residuals) RSquared() float64 {
	average := ss.Sum / float64(ss.Count)
	var ssTot float64
	for _, y := range ss.ys {
		tot := y - average
		ssTot += tot * tot
	}
	return 1.0 - (ss.SSRes / ssTot)
}

func (ss *Residuals) MSE() float64 {
	return ss.SSRes / float64(ss.Count)
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
	ss.SSRes += e * e
}
