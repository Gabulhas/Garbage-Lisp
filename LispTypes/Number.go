package LispTypes

import "fmt"

type Number struct {
	Contents float32
}

func (number Number) GetType() InterfaceType {
	return NUMBER
}

func (number Number) GetContent() float32 {
	return number.Contents
}

func (number Number) ToString() string {
	return fmt.Sprintf("%s %.2f", number.GetType().ToString(), number.GetContent())
}
