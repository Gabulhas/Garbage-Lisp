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

func (Procedure) InitLambda(arguments, lambdaContent LispToken) Procedure {

	var argumentsAsString []string

	expContent, isExp := arguments.(List)
	if !isExp {
		log.Fatal("::ERROR:: Lambda argument should be expression. (arg1, arg2, ...).")
	}

	argumentList := expContent.GetContent()

	for _, argument := range argumentList {
		if value, ok := argument.(Symbol); ok {
			argumentsAsString = append(argumentsAsString, value.GetContent())
		} else {
			log.Fatal("::ERROR:: Lambda arguments can only be symbols.")
		}
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

func (procedure Procedure) ToString() string {
	if procedure.Native {
		return fmt.Sprintf("PROCEDURE:NATIVE %s", procedure.Name)
	} else {
		return fmt.Sprintf("PROCEDURE:USER %s %s -> %s", procedure.Name, procedure.Arguments, procedure.LambdaContent.ValueToString())
	}
}

func (procedure Procedure) ValueToString() string {
	if procedure.Native {
		return fmt.Sprintf("PROCEDURE %s", procedure.Name)
	} else {
		return fmt.Sprintf("PROCEDURE %s", procedure.Name)
	}
}
