package LispTypes

import (
	"fmt"
	"log"
)

type ProcedureFunction func(tokens ...LispToken) LispToken

type Procedure struct {
	Name          string
	Native        bool
	NativeContent ProcedureFunction
	LambdaContent LispToken
	Arguments     []string
}

func (procedure Procedure) GetType() InterfaceType {
	return PROCEDURE
}

func (procedure Procedure) ToString() string {
	return fmt.Sprintf("%s %s %s", procedure.GetType().ToString(), procedure.LambdaContent.ToString(), procedure.Name)
}

func (Procedure) InitLambda(arguments, lambdaContent LispToken) Procedure {

	var argumentsAsString []string

	expContent, isExp := arguments.(Exp)
	if !isExp {
		log.Fatal("::ERROR:: Lambda argument should be expression. (arg1, arg2, ...).")
	}

	if argumentList, isList := expContent.GetContent().(List); isList {

		for _, argument := range argumentList.GetContent() {
			if value, ok := argument.(Symbol); ok {
				argumentsAsString = append(argumentsAsString, value.GetContent())
			} else {
				log.Fatal("::ERROR:: Lambda arguments can only be symbols.")
			}
		}

	} else {
		log.Fatal("::ERROR:: Lambda arguments should be in list (arg1, arg2, ...).")
	}
	return Procedure{
		Name:          "Anonymous",
		Native:        false,
		NativeContent: nil,
		LambdaContent: lambdaContent,
		Arguments:     argumentsAsString,
	}
}
func (procedure Procedure) Call(env map[string]LispToken, params ...LispToken) LispToken {
	if procedure.Native {
		return procedure.NativeContent(params...)
	} else {
		if len(params) != len(procedure.Arguments) {
			log.Fatalf("\n::ERROR:: Procedure [%s] arguments unmatched: %d =/= %d.", procedure.Name, len(params), len(procedure.Arguments))
		}

		for i, argName := range procedure.Arguments {
			env[argName] = params[i]
		}
		return procedure.LambdaContent
	}
}

func (procedure Procedure) ValueToString() string {
	return fmt.Sprintf("PROCEDURE %s", procedure.Name)
}
