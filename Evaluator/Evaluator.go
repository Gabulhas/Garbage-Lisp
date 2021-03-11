package Evaluator

import (
	"GarbageLisp/Env"
	"GarbageLisp/LispTypes"
)

type Evaluator struct {
	currentEnv Env.Env  //Current or Outer Env
	innerEnv   *Env.Env //Inner Env
}

func NewEval() *Evaluator {
	neweval := new(Evaluator)
	neweval.currentEnv = Env.InitStandardEnv()
	innerEnv := new(Env.Env)
	innerEnv.Contents = map[string]LispTypes.LispToken{}
	innerEnv.Using = false
	neweval.innerEnv = innerEnv
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

func (evaluator Evaluator) FindEnv(key string) (Env.Env, bool) {
	if _, isInEnv := evaluator.innerEnv.Contents[key]; isInEnv {
		return *evaluator.innerEnv, true
	}
	if _, isInEnv := evaluator.currentEnv.Contents[key]; isInEnv {
		return evaluator.currentEnv, true
	}
	return Env.Env{}, false
}

func (evaluator Evaluator) Define(key string, token LispTypes.LispToken) {
	if evaluator.innerEnv.Using {
		evaluator.innerEnv.Contents[key] = token
		return
	}
	evaluator.currentEnv.Contents[key] = token

}
