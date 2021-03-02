package LispTypes

import (
	"fmt"
	"strings"
)

type Env struct {
	Contents map[string]LispToken
}

func (env Env) GetType() InterfaceType {
	return EXP
}

func (env Env) GetContent() map[string]LispToken {
	return env.Contents
}

func (env Env) ToString() string {
	var b strings.Builder
	for key, p := range env.Contents {
		fmt.Fprintf(&b, "[%s : %s]", key, p.ToString())
	}
	return b.String()
}

// func GetStandardEnv