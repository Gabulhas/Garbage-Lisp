package Evaluator

import (
	"GarbageLisp/LispTypes"
	"log"
)


//TODO: mudar de return LispToken To Exp
func (evaluator *Evaluator) Run(parsedTokens LispTypes.LispToken) LispTypes.LispToken {

	switch value := parsedTokens.(type) {

	case LispTypes.Symbol:
		return evaluator.currentEnv.Contents[value.Contents]
		break
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
		log.Fatal("Expression Not Starting With Symbol")
	}
	// "Builtins"
	//TODO: Change to swtich case
	if symbol == "define" {
		newVariableName, err := LispTypes.GetSymbolContent(content[1])
		if err != nil {
			log.Fatal("Variable Name Not A Symbol")
		}

		evaluator.currentEnv.Contents[newVariableName] = evaluator.Run(content[len(content)-1])
		return nil
	} else if symbol == "if" {
		if len(content) != 4 {
			return nil
		}
		//if
		test := content[1]
		//then
		conseq := content[2]
		//else
		alt := content[3]

		if value, ok := evaluator.Run(test).(LispTypes.LispBoolean); ok {
			if value.Contents {
				return evaluator.Run(conseq)
			}
			return evaluator.Run(alt)

		} else {
			log.Fatal("If Condition Is Not A Boolean")
		}

	} else if symbol == "lambda" {

	} else if symbol == "quote" {
		return content[1]
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
			return resultFunc.Call(nil, arguments...)
			break
		default:
			log.Fatal("Not a procedure")

		}
		return nil

	}

	log.Printf("My Env")
	log.Println(evaluator.currentEnv.Contents)
	log.Printf("\n\n\nFailed SExpression")
	log.Printf("%s", list.ToString())
	log.Fatal("Error, unexpected type")
	return nil
}
