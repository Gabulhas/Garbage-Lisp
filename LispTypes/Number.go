package LispTypes

import "fmt"

//TODO: Encapsulating Float or Int
type Number struct {
	IsFloat       bool
	IntContents   int32
	FloatContents float64
}

func (num Number) GetType() InterfaceType {
	return NUMBER
}

func (num Number) GetContent() (int32, float64) {
	if num.IsFloat {
		return 0, num.FloatContents
	} else {
		return num.IntContents, 0
	}
}

func (num Number) ToString() string {
	if num.IsFloat {
		return fmt.Sprintf("%s %.2f", num.GetType().ToString(), num.FloatContents)
	} else {
		return fmt.Sprintf("%s %.2f", num.GetType().ToString(), num.FloatContents)
	}
}

func (num Number) ValueToString() string {
	if num.IsFloat {
		return fmt.Sprintf("%.6f", num.FloatContents)
	} else {
		return fmt.Sprintf("%d", num.IntContents)
	}
}

func NewInt(val int32) Number {
	return Number{IsFloat: false, IntContents: val}
}

func NewFloat(val float64) Number {
	return Number{IsFloat: true, FloatContents: val}
}
