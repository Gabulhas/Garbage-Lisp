package Env

import "GarbageLisp/LispTypes"

func unpackList(token LispTypes.LispToken) []LispTypes.LispToken {

	switch v := token.(type) {
	case LispTypes.List:
		return v.Contents
		break
	default:
		return nil
	}
	return nil
}
