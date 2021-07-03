package Env

import "github.com/Gabulhas/Garbage-Lisp/LispTypes"

func unpackList(token LispTypes.LispToken) []LispTypes.LispToken {

	switch v := token.(type) {
	case LispTypes.List:
		return v.Contents
	default:
		return nil
	}
}
