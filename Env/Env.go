package Env

import (
	"GarbageLisp/LispTypes"
	"math"
)

type Env struct {
	Contents map[string]LispTypes.LispToken
	Using    bool
}

func InitStandardEnv() Env {
	//TODO: Move different file
	envmap := map[string]LispTypes.LispToken{
		"pi":         NumberFromConstant(math.Pi),
		"e":          NumberFromConstant(math.E),
		"max_number": NumberFromConstant(math.MaxFloat64),
	}
	newEnv := Env{Contents: envmap, Using: true}
	InitEnvNativeFunctions(newEnv)
	return newEnv
}

func NumberFromConstant(number float64) LispTypes.Number {
	return LispTypes.Number{Contents: number}
}

func (env Env) AddProcedureFromFunction(procedureFunction LispTypes.ProcedureFunction, name string) {
	env.Contents[name] = LispTypes.Procedure{
		Name:          name,
		Native:        true,
		NativeContent: procedureFunction,
		LambdaContent: nil,
	}
}
