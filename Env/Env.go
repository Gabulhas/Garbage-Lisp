package Env

import (
	"GarbageLisp/LispTypes"
	"math"
)

type Env struct {
	Contents map[string]LispTypes.LispToken
}

func InitStandardEnv() Env {
	envmap := map[string]LispTypes.LispToken{
		"pi": NumberFromConstant(math.Pi),
		"e":  NumberFromConstant(math.E),
	}
	InitEnvNativeFunctions(envmap)
	return Env{Contents: envmap}
}


func NumberFromConstant(number float64) LispTypes.Number {
	return LispTypes.Number{Contents: number}
}

func ProcedureFromFunction(procedureFunction LispTypes.ProcedureFunction) LispTypes.Procedure {
	return LispTypes.Procedure{
		Native:        true,
		NativeContent: procedureFunction,
		LambdaContent: nil,
	}
}
