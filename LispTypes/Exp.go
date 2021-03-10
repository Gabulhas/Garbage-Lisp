package LispTypes

import (
	"errors"
	"fmt"
)

type Exp struct {
	Contents LispToken
}

func (exp Exp) GetType() InterfaceType {
	return EXP
}

func (exp Exp) GetContent() LispToken {
	return exp.Contents
}

func (exp Exp) ToString() string {
	return fmt.Sprintf("%s (%s)", exp.GetType().ToString(), exp.Contents.ToString())
}

func (exp *Exp) ContainsList() bool {
	if _, ok := exp.Contents.(List); ok {
		return true
	}
	return false
}

func (exp *Exp) GetList() (List, error) {
	if list, ok := exp.Contents.(List); ok {
		return list, nil
	}
	return List{}, errors.New("NotAList")
}

func (exp *Exp) AppendIfList(token LispToken) error {

	switch v := exp.Contents.(type) {
	case List:
		exp.Contents = v.Append(token)
		return nil
	default:
		return errors.New("WrongType")
	}
}

func (exp *Exp) ChangeIfAtom(token LispToken) error {

	switch content := exp.Contents.(type) {

	case Atom:
		switch tokenType := token.(type) {
		case Symbol, Number:
			exp.Contents = Atom{Contents: tokenType}
		default:
			return errors.New("WrongAtomType" + content.ToString())

		}
	default:
		return errors.New("WrongAppendingToAtom")
	}

	return nil
}

func (exp Exp) ValueToString() string {
	return fmt.Sprintf("%s", exp.Contents.ValueToString())
}
