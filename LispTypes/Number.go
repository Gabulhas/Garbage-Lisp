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
		return fmt.Sprintf("%s:FLOAT %f", num.GetType().ToString(), num.FloatContents)
	} else {
		return fmt.Sprintf("%s:INT %d", num.GetType().ToString(), num.IntContents)
	}
}

func (num Number) ValueToString() string {
	if num.IsFloat {
		return fmt.Sprintf("%.6f", num.FloatContents)
	} else {
		return fmt.Sprintf("%d", num.IntContents)
	}
}

func (num Number) GetAsFloat() float64 {
	return float64(num.IntContents) + num.FloatContents
}

func (num Number) GetAsInt() int32 {
	return num.IntContents + int32(num.FloatContents)
}

func NewInt(val int32) Number {
	return Number{IsFloat: false, IntContents: val}
}

func NewFloat(val float64) Number {
	return Number{IsFloat: true, FloatContents: val}
}

/*
################################
TODO: CLEAN THIS GARBAGE OMG.

################################
*/
func Add(a, b Number) Number {
	aInt, aFloat := a.GetContent()
	bInt, bFloat := b.GetContent()
	if !a.IsFloat && !b.IsFloat {
		return NewInt(aInt + bInt)
	}
	return NewFloat((float64(aInt) + aFloat) + (float64(bInt) + bFloat))
}

func Sub(a, b Number) Number {
	aInt, aFloat := a.GetContent()
	bInt, bFloat := b.GetContent()
	if !a.IsFloat && !b.IsFloat {
		return NewInt(aInt - bInt)
	}
	return NewFloat((float64(aInt) + aFloat) - (float64(bInt) + bFloat))
}

func Div(a, b Number) Number {
	aInt, aFloat := a.GetContent()
	bInt, bFloat := b.GetContent()
	if !a.IsFloat && !b.IsFloat {
		return NewInt(aInt / bInt)
	}
	return NewFloat((float64(aInt) + aFloat) / (float64(bInt) + bFloat))
}

func Mul(a, b Number) Number {
	aInt, aFloat := a.GetContent()
	bInt, bFloat := b.GetContent()
	if !a.IsFloat && !b.IsFloat {
		return NewInt(aInt * bInt)
	}
	return NewFloat((float64(aInt) + aFloat) * (float64(bInt) + bFloat))
}
