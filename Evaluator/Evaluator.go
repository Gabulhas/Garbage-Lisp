package Evaluator

import (
	"GarbageLisp/Env"
	"GarbageLisp/LispTypes"
)

type Evaluator struct {
	//Outer -> inner envs
	envs []*Env.Env
}

func NewEval() *Evaluator {
	neweval := new(Evaluator)
	outer_env := new(Env.Env)
	outer_env.Contents = Env.InitStandardEnv().Contents
	neweval.envs = append(neweval.envs, outer_env)
	return neweval
}

func (evaluator Evaluator) FindValue(key string) LispTypes.LispToken {

	for i := len(evaluator.envs) - 1; i >= 0; i-- {
		if value, isInEnv := evaluator.envs[i].Contents[key]; isInEnv {
			return value
		}
	}
	//We return false which is null
	return LispTypes.LispBoolean{Contents: false}
}

func (evaluator Evaluator) FindEnv(key string) (Env.Env, bool) {
	for i := len(evaluator.envs) - 1; i >= 0; i-- {
		if _, isInEnv := evaluator.envs[i].Contents[key]; isInEnv {
			return *(evaluator.envs[i]), true
		}
	}
	return Env.Env{}, false
}

func (evaluator Evaluator) Define(key string, token LispTypes.LispToken) {
	evaluator.envs[len(evaluator.envs)-1].Contents[key] = token

}
