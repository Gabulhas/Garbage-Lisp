package Evaluator

import (
	"GarbageLisp/Env"
	"GarbageLisp/LispTypes"
)

type Evaluator struct {
	currentEnv Env.Env //Current or Outer Env
	innerEnv   Env.Env //Inner Env
}

func NewEval() *Evaluator {
	neweval := new(Evaluator)
	neweval.currentEnv = Env.InitStandardEnv()
	neweval.innerEnv = Env.Env{Contents: map[string]LispTypes.LispToken{}}
	return neweval
}

func (evaluator Evaluator) FindValue(key string) LispTypes.LispToken {

	if value, isInList := evaluator.innerEnv.Contents[key]; isInList {
		return value
	}
	if value, isInList := evaluator.currentEnv.Contents[key]; isInList {
		return value
	}

	//We return false which is null
	return LispTypes.LispBoolean{Contents: false}
}
