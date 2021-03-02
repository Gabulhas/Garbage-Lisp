package LispTypes

import "errors"

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
	return exp.Contents.ToString()
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
