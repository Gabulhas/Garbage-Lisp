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

func (exp *Exp) AppendIfList(token LispToken) error {

	switch v := exp.Contents.(type) {
	case List:
		exp.Contents = v.Append(token)
		return nil
	default:
		return errors.New("WrongType")
	}
}

func (exp Exp) ValueToString() string {
	return exp.Contents.ValueToString()
}
