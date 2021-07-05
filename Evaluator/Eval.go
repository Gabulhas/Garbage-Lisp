package Evaluator

import (
	"log"
	"strings"

	"github.com/Gabulhas/Garbage-Lisp/Env"
	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
	"github.com/Gabulhas/Garbage-Lisp/Parser"
)

func (evaluator *Evaluator) Run(parsedTokens LispTypes.LispToken) LispTypes.LispToken {

	switch value := parsedTokens.(type) {

	case LispTypes.Symbol:
		return evaluator.FindValue(value.Contents)
	case LispTypes.List:
		return evaluator.evalSEXPRESSION(value)
	case LispTypes.Exp:
		return evaluator.Run(value.GetContent())
	default:
		return parsedTokens
	}
}

func (evaluator *Evaluator) evalSEXPRESSION(list LispTypes.List) LispTypes.LispToken {
	content := list.Contents
	symbol, err := LispTypes.GetSymbolContent(content[0])
	if err != nil {
		log.Printf("::ERROR:: %s Expression Not Starting With Symbol", list.ToString())
		return nil

	}
	// "Builtins"
	if strings.EqualFold(symbol, "define") {

		var newVariableName string

		//TODO: Change this garbage
		if content[1].GetType() == LispTypes.STRING || content[1].GetType() == LispTypes.SYMBOL {
			newVariableName = content[1].ValueToString()
		} else if content[1].GetType() == LispTypes.LIST || content[1].GetType() == LispTypes.EXP {
			resultingSymbol := evaluator.Run(content[1])
			if resultingSymbol.GetType() == LispTypes.STRING || resultingSymbol.GetType() == LispTypes.SYMBOL {
				newVariableName = resultingSymbol.ValueToString()
			} else {
				log.Printf("\n::ERROR:: %s Not A Symbol.", content[1].ToString())
				return nil
			}

		} else {
			log.Printf("\n::ERROR:: %s Not A Symbol.", content[1].ToString())
			return nil
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

		if value, ok := test.(LispTypes.LispBoolean); ok {
			testResult = value.GetContent()
		} else {
			log.Printf("\n::ERROR:: if first argument must be a boolean: if (boolean) (then case) (else case). Got %s", test.ToString())
			return nil
		}

		if testResult {
			return evaluator.Run(conseq)
		}
		return evaluator.Run(alt)

	} else if strings.EqualFold(symbol, "lambda") {

		if len(content) != 3 {
			log.Printf("\n::ERROR:: Lambda requires 2 expressions: lambda (arguments) (body)\n Got %d instead", len(content))
			return nil

		}
		return LispTypes.Procedure.InitLambda(LispTypes.Procedure{}, content[1], content[2])

	} else if strings.EqualFold(symbol, "map") {
		if len(content) != 3 {
			log.Printf("::ERROR:: map requires 2 expressions: map (procedure) (list)")
			return nil
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
			expression = evaluator.Run(value)
		} else {
			expression = content[1]
		}

		return evaluator.Run(expression)

	} else if strings.EqualFold(symbol, "set!") {
		newVariableName, err := LispTypes.GetSymbolContent(content[1])
		if err != nil {
			log.Printf("\n::ERROR:: %s Not a Symbol.", content[1])
			return nil
		}
		env, exists := evaluator.FindEnv(newVariableName)
		if !exists {
			log.Println("::ERROR:: Cannot set! a non-existing symbol")
			return nil
		}
		env.Contents[newVariableName] = evaluator.Run(content[len(content)-1])
		return nil

	} else if strings.EqualFold(symbol, "load") {
		if len(content) != 2 {
			log.Println("::ERROR:: Cannot load another script. Usage: load \"Path To File\"")
		} else if filename, ok := content[1].(LispTypes.LispString); !ok {
			log.Println("::ERROR:: Cannot load another script. Usage: load \"Path To File\"")
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
			if resultFunc.Native {
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
			log.Printf("\n::ERROR:: %s: Not a procedure.", symbol)

		}
		return nil

	}
}
