package LispTypes

import "fmt"

type LispString struct {
	Contents string
}

func (lispString LispString) GetType() InterfaceType {
	return STRING
}

func (lispString LispString) GetContent() string {
	return lispString.Contents
}

func (lispString LispString) ToString() string {
	return fmt.Sprintf("%s %s", lispString.GetType().ToString(), lispString.GetContent())
}

func (lispString LispString) ValueToString() string {
	return lispString.Contents
}
