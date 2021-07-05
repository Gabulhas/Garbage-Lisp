package Env

import (
	"math"

	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
)

func InitEnvNativeConstants(env Env) {

	env.Contents["pi"] = NumberFromConstant(math.Pi)
	env.Contents["e"] = NumberFromConstant(math.E)
	env.Contents["max_number"] = NumberFromConstant(math.MaxFloat64)

}

func NumberFromConstant(number float64) LispTypes.Number {
	return LispTypes.NewFloat(number)
}
