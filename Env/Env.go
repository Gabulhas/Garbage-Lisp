package Env

import (
	"GarbageLisp/LispTypes"
)

type Env struct {
	Contents map[string]LispTypes.LispToken
	Using    bool
}

func InitStandardEnv() Env {
	envmap := map[string]LispTypes.LispToken{}
	newEnv := Env{Contents: envmap, Using: true}
	InitEnvNativeConstants(newEnv)
	InitEnvNativeFunctions(newEnv)
	return newEnv
}

func (env Env) AddProcedureFromFunction(procedureFunction LispTypes.ProcedureFunction, name string) {
	env.Contents[name] = LispTypes.Procedure{
		Name:          name,
		Native:        true,
		NativeContent: procedureFunction,
		LambdaContent: nil,
	}
}
