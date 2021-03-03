package LispTypes

import "fmt"

type Number struct {
	Contents float64
}

func (number Number) GetType() InterfaceType {
	return NUMBER
}

func (number Number) GetContent() float64 {
	return number.Contents
}

func (number Number) ToString() string {
	return fmt.Sprintf("%s %.2f", number.GetType().ToString(), number.GetContent())
}

func ValueToNumber(value float64) Number {
	return Number{Contents: value}
}
func (number Number) ValueToString() string {
	return fmt.Sprintf("%.3f", number.Contents)
}
