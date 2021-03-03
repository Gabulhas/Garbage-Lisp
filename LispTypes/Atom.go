package LispTypes

import (
	"errors"
	"fmt"
)

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

func TokenToAtom(token LispToken) (Atom, error) {
	if IsNumber(token) || IsSymbol(token) {
		return Atom{Contents: token}, nil
	} else {
		return Atom{}, errors.New("NoPossibleToken")
	}
}

func (atom Atom) ValueToString() string {
	return fmt.Sprintf("%s", atom.ValueToString())
}
