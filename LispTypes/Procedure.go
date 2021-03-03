package LispTypes

import "fmt"

type ProcedureFunction func(tokens ...LispToken) LispToken

type Procedure struct {
	Native        bool
	NativeContent ProcedureFunction
	LambdaContent LispToken
	ProcedureEnv  map[string]LispToken
}

func (procedure Procedure) GetType() InterfaceType {
	return PROCEDURE
}

func (procedure Procedure) ToString() string {
	return fmt.Sprintf("%s %s", procedure.GetType().ToString(), procedure.LambdaContent.ToString())
}

func (Procedure) InitLambda(lambdaContent LispToken) Procedure {
	return Procedure{
		Native:        false,
		NativeContent: nil,
		LambdaContent: lambdaContent,
	}
}

func (procedure Procedure) Call(env map[string]LispToken, params ...LispToken) LispToken {
	if procedure.Native {
		return procedure.NativeContent(params...)
	}
	return nil
}

//TODO: something to print function names (?)
func (procedure Procedure)ValueToString() string {
	return fmt.Sprintf("FUNCTION")
}
