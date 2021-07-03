package LispTypes

import (
	"errors"
)

func GetNumberContent(token LispToken) (float64, error) {
	switch result := token.(type) {
	case Number:
		return result.GetContent(), nil
	case LispBoolean:
		if result.GetContent() {
			return 1.0, nil
		} else {
			return 0.0, nil
		}
	default:
		return -1, errors.New("WrongType")
	}
}

func GetSymbolContent(token LispToken) (string, error) {
	switch result := token.(type) {
	case Symbol:
		return result.GetContent(), nil
	default:
		return "", errors.New("WrongType")
	}
}
