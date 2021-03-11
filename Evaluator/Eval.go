package Evaluator

import (
	"GarbageLisp/Env"
	"GarbageLisp/LispTypes"
	"log"
	"strings"
)

//TODO: mudar de return LispToken To Exp
func (evaluator *Evaluator) Run(parsedTokens LispTypes.LispToken) LispTypes.LispToken {

	switch value := parsedTokens.(type) {

	case LispTypes.Symbol:
		return evaluator.FindValue(value.Contents)
	case LispTypes.Number:
		return parsedTokens
	case LispTypes.List:
		return evaluator.evalS_Expression(value)
	case LispTypes.LispBoolean:
		return parsedTokens
	case LispTypes.Exp:
		return evaluator.Run(value.GetContent())
	}
	log.Fatal("Error, unexpected type A")
	return nil
}

func (evaluator *Evaluator) evalS_Expression(list LispTypes.List) LispTypes.LispToken {
	//log.Println(list.ToString())
	content := list.Contents
	symbol, err := LispTypes.GetSymbolContent(content[0])
	if err != nil {
		log.Fatalf("%s :Expression Not Starting With Symbol", list.ToString())
	}
	// "Builtins"
	//TODO: Change to swtich case
	if strings.EqualFold(symbol, "define") {
		newVariableName, err := LispTypes.GetSymbolContent(content[1])
		if err != nil {
			log.Fatal("Variable Name Not A Symbol")
		}

		evaluator.Define(newVariableName, evaluator.Run(content[len(content)-1]))
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
			if value.GetContent() != 0 {
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
			log.Fatal("Lambda requires 2 elements, lambda (arguments) (body)")
		}
		return LispTypes.Procedure.InitLambda(LispTypes.Procedure{}, content[1], content[2])

	} else if strings.EqualFold(symbol, "all") {
		for _, exp := range content {
			evaluator.Run(exp)
		}
		return nil
	} else if strings.EqualFold(symbol, "quote") {
		return content[1]

	} else if strings.EqualFold(symbol, "set!") {
		newVariableName, err := LispTypes.GetSymbolContent(content[1])
		if err != nil {
			log.Fatal("Variable Name Not A Symbol")
		}
		env, exists := evaluator.FindEnv(newVariableName)
		if !exists {
			log.Fatal("Cannot !set an non-existing symbol")
		}
		env.Contents[newVariableName] = evaluator.Run(content[len(content)-1])
		return nil
	} else {
		//Todo: alterar
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
				innerEnv.Using = true
				lambdaBody := resultFunc.Call(innerEnv.Contents, arguments...)
				newEvaluator := Evaluator{}
				newEvaluator.currentEnv = evaluator.currentEnv
				newEvaluator.innerEnv = innerEnv
				return newEvaluator.Run(lambdaBody)
			}
		default:
			log.Fatalf("%s: Not a procedure", symbol)

		}
		return nil

	}

	log.Println(evaluator.currentEnv.Contents)
	log.Printf("\n\n\nFailed SExpression")
	log.Printf("%s", list.ToString())
	log.Fatal("Error, unexpected type")
	return nil
}
