package Evaluator

import (
	"github.com/Gabulhas/Garbage-Lisp/Env"
	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
	"github.com/Gabulhas/Garbage-Lisp/Parser"
	"log"
	"strings"
)

func (evaluator *Evaluator) Run(parsedTokens LispTypes.LispToken) LispTypes.LispToken {

	switch value := parsedTokens.(type) {

	case LispTypes.Symbol:
		return evaluator.FindValue(value.Contents)
	case LispTypes.LispString:
		return parsedTokens
	case LispTypes.Number:
		return parsedTokens
	case LispTypes.List:
		return evaluator.evalS_Expression(value)
	case LispTypes.LispBoolean:
		return parsedTokens
	case LispTypes.Exp:
		return evaluator.Run(value.GetContent())
	}
	log.Printf("\n::ERROR:: Unexpected Type")
	return nil
}

func (evaluator *Evaluator) evalS_Expression(list LispTypes.List) LispTypes.LispToken {
	content := list.Contents
	symbol, err := LispTypes.GetSymbolContent(content[0])
	if err != nil {
		log.Fatalf("\n::ERROR:: %s Expression Not Starting With Symbol", list.ToString())
	}
	// "Builtins"
	if strings.EqualFold(symbol, "define") {

		var newVariableName string

		//TODO: Change this garbage
		if content[1].GetType() == LispTypes.STRING || content[1].GetType() == LispTypes.SYMBOL {
			newVariableName = content[1].ValueToString()
		} else if content[1].GetType() == LispTypes.LIST || content[1].GetType() == LispTypes.EXP {
			resultingSymbol := evaluator.Run(LispTypes.UnpackFromExp(content[1]))
			if resultingSymbol.GetType() == LispTypes.STRING || resultingSymbol.GetType() == LispTypes.SYMBOL {
				newVariableName = resultingSymbol.ValueToString()
			} else {
				log.Fatalf("\n::ERROR:: %s Not A Symbol.", content[1])
			}

		} else {
			log.Fatalf("\n::ERROR:: %s Not A Symbol.", content[1].ToString())
		}

		evaluatedDefine := evaluator.Run(content[len(content)-1])
		if value, ok := evaluatedDefine.(LispTypes.Procedure); ok {
			temp := value
			temp.Name = newVariableName
			evaluator.Define(newVariableName, temp)
			return nil
		}

		evaluator.Define(newVariableName, evaluatedDefine)
		return nil

	} else if strings.EqualFold(symbol, "if") {
		if len(content) != 4 {
			return nil
		}
		//if
		test := evaluator.Run(content[1])
		//then
		conseq := content[2]
		//else
		alt := content[3]

		var testResult bool
		switch value := test.(type) {
		case LispTypes.LispBoolean:
			testResult = value.GetContent()
			break
		case LispTypes.List:
			if len(list.GetContent()) > 0 {
				testResult = true
			} else {
				testResult = false
			}
			break
		case LispTypes.Number:
			if value.GetContent() > 0 {
				testResult = true
			} else {
				testResult = false
			}
			break
		}

		if testResult {
			return evaluator.Run(conseq)
		}
		return evaluator.Run(alt)

	} else if strings.EqualFold(symbol, "lambda") {

		if len(content) != 3 {
			log.Fatalf("\n::ERROR:: Lambda requires 2 expressions: lambda (arguments) (body)\n Got %d instead", len(content))

		}
		return LispTypes.Procedure.InitLambda(LispTypes.Procedure{}, content[1], content[2])

	} else if strings.EqualFold(symbol, "map") {
		if len(content) != 3 {
			log.Fatal("::ERROR:: map requires 2 expressions: map (procedure) (list)")
		}
		var result []LispTypes.LispToken

		arguments := evaluator.Run(content[2])
		procedure := content[1]

		for _, token := range LispTypes.Unpack(arguments) {

			newExp := LispTypes.List{Contents: []LispTypes.LispToken{procedure, token}}
			tokenResult := evaluator.Run(newExp)

			result = append(result, tokenResult)
		}
		return LispTypes.List{Contents: result}

	} else if strings.EqualFold(symbol, "all") {
		for _, exp := range content {
			evaluator.Run(exp)
		}
		return nil
	} else if strings.EqualFold(symbol, "quote") {
		if value, ok := content[1].(LispTypes.Exp); ok {
			return value.Contents

		} else {
			return content[1]
		}

	} else if strings.EqualFold(symbol, "eval") {
		var expression LispTypes.LispToken
		if value, ok := content[1].(LispTypes.Symbol); ok {
			symbol_value := evaluator.Run(value)
			expression = LispTypes.UnpackFromExp(symbol_value)

		} else {
			expression = LispTypes.UnpackFromExp(content[1])
		}

		return evaluator.Run(expression)

	} else if strings.EqualFold(symbol, "set!") {
		newVariableName, err := LispTypes.GetSymbolContent(content[1])
		if err != nil {
			log.Fatalf("\n::ERROR:: %s Not a Symbol.", content[1])
		}
		env, exists := evaluator.FindEnv(newVariableName)
		if !exists {
			log.Fatal("::ERROR:: Cannot set! a non-existing symbol")
		}
		env.Contents[newVariableName] = evaluator.Run(content[len(content)-1])
		return nil

	} else if strings.EqualFold(symbol, "load") {
		if len(content) != 2 {
			log.Fatal("::ERROR:: Cannot load another script. Usage: load \"Path To File\"")
		}
		if filename, ok := content[1].(LispTypes.LispString); !ok {
			log.Fatal("::ERROR:: Cannot load another script. Usage: load \"Path To File\"")
		} else {
			parsed := Parser.ParseFromFile(filename.Contents)
			evaluator.Run(parsed)
		}

		return nil

	} else {
		var arguments []LispTypes.LispToken
		for i, args := range content {
			if i == 0 {
				continue
			}
			arguments = append(arguments, evaluator.Run(args))
		}

		switch resultFunc := evaluator.Run(content[0]).(type) {
		case LispTypes.Procedure:
			if resultFunc.IsNative() {
				return resultFunc.Call(nil, arguments...)
			} else {

				innerEnv := new(Env.Env)
				innerEnv.Contents = make(map[string]LispTypes.LispToken)

				lambdaBody := resultFunc.Call(innerEnv.Contents, arguments...)
				newEvaluator := new(Evaluator)

				for i := 0; i < len(evaluator.envs); i++ {
					newEvaluator.envs = append(newEvaluator.envs, evaluator.envs[i])
				}
				newEvaluator.envs = append(newEvaluator.envs, innerEnv)

				return newEvaluator.Run(lambdaBody)
			}
		default:
			log.Fatalf("\n::ERROR:: %s: Not a procedure.", symbol)

		}
		return nil

	}
}
