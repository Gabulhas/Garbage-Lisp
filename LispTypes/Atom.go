package LispTypes

import "fmt"

type Atom struct {
	Contents LispToken
}

func (atom Atom) GetType() InterfaceType {
	return ATOM
}

func (atom Atom) GetContent() LispToken {
	return atom.Contents
}

func (atom Atom) ToString() string {
	return fmt.Sprintf("%s (%s)", atom.GetType().ToString(), atom.Contents.ToString())
}
