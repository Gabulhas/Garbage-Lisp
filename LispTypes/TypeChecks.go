package LispTypes

import (
	"errors"
)

func IsSymbol(token LispToken) bool {
	return token.GetType() == SYMBOL
}
func IsNumber(token LispToken) bool {
	return token.GetType() == NUMBER
}
func IsAtom(token LispToken) bool {
	tokenType := token.GetType()
	return tokenType == ATOM || tokenType == NUMBER || tokenType == SYMBOL
}
func IsList(token LispToken) bool {
	return token.GetType() == LIST
}
func IsEXP(token LispToken) bool {
	return token.GetType() == EXP
}
func IsProcedure(token LispToken) bool {
	return token.GetType() == PROCEDURE
}

func GetNumberContent(token LispToken) (float64, error) {
	switch result := token.(type) {
	case Number:
		return result.GetContent(), nil
		break
	case Atom:
		return GetNumberContent(result)
		break
	default:
		return -1, errors.New("WrongType")
	}
	return -1, errors.New("WrongType")
}

func GetSymbolContent(token LispToken) (string, error) {
	switch result := token.(type) {
	case Symbol:
		return result.GetContent(), nil
		break
	case Atom:
		return GetSymbolContent(result)
		break
	default:
		return "", errors.New("WrongType")
	}
	return "", errors.New("WrongType")
}

func GetProcedureCall(token LispToken) (func(env map[string]LispToken, params ...LispToken) LispToken, error) {
	if procedure, ok := token.(Procedure); ok {
		return procedure.Call, nil
	} else {
		return nil, errors.New("NotAFunction")
	}

}
