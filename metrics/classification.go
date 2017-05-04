package metrics

type ConfusionMatrix struct {
	TP float64
	FP float64
	TN float64
	FN float64
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
