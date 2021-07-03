package Env

import (
	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
	"os"
)

type Env struct {
	Contents map[string]LispTypes.LispToken
}

func InitStandardEnv() Env {
	envmap := map[string]LispTypes.LispToken{}
	newEnv := Env{Contents: envmap}
	InitEnvNativeConstants(newEnv)
	InitEnvNativeFunctions(newEnv)
	FilterAndAddCommandLineArgs(newEnv)

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

func FilterAndAddCommandLineArgs(env Env) {
	currentArgs := os.Args[1:]
	var argsAsStrings []LispTypes.LispToken
	if len(currentArgs) > 1 {
		if currentArgs[0] == "-load" {
			currentArgs = currentArgs[1:]
		}
	}

	for _, arg := range currentArgs {
		argsAsStrings = append(argsAsStrings, LispTypes.LispString{Contents: arg})
	}

	env.Contents["args"] = LispTypes.List{Contents: argsAsStrings}
}
