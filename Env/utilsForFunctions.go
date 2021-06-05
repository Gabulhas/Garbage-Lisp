package Env

import "GarbageLisp/LispTypes"

func unpackList(token LispTypes.LispToken) []LispTypes.LispToken {

	switch v := token.(type) {
	case LispTypes.List:
		return v.Contents
	default:
		return nil
	}
}
