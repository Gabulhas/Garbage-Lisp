package LispTypes

import "fmt"

type Symbol struct {
	Contents string
}

func (symbol Symbol) GetType() InterfaceType {
	return SYMBOL
}

func (symbol Symbol) GetContent() string {
	return symbol.Contents
}

func (symbol Symbol) ToString() string {
	return fmt.Sprintf("%s %s", symbol.GetType().ToString(), symbol.GetContent())
}

func (symbol Symbol) ValueToString() string {
	return fmt.Sprintf("%s", symbol.Contents)
}
