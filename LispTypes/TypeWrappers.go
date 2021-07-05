package LispTypes

import (
	"errors"
)

func GetNumberContent(token LispToken) (int32, float64, error) {
	if result, ok := token.(Number); ok {
		intVal, floatVal := result.GetContent()
		return intVal, floatVal, nil
	} else {
		return 0, 0, errors.New("WrongType")
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

func GetNumberAsFloat(token LispToken) (float64, error) {
	if result, ok := token.(Number); ok {
		return result.GetAsFloat(), nil
	} else {
		return 0, errors.New("WrongType")
	}
}
