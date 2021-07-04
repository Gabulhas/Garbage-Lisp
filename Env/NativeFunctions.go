package Env

import (
	"fmt"
	"log"
	"strings"

	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
)

func InitEnvNativeFunctions(env Env) {
	//Arithmetic
	env.AddProcedureFromFunction(multiply, "*")
	env.AddProcedureFromFunction(add, "+")
	env.AddProcedureFromFunction(sub, "-")
	env.AddProcedureFromFunction(divide, "/")
	env.AddProcedureFromFunction(modulo, "%")

	env.AddProcedureFromFunction(maxnumber, "max")
	env.AddProcedureFromFunction(minnumber, "min")

	env.AddProcedureFromFunction(intpart, "intPart")

	//Other
	env.AddProcedureFromFunction(begin, "begin")
	env.AddProcedureFromFunction(printLisp, "print")
	env.AddProcedureFromFunction(printfLisp, "printf")
	env.AddProcedureFromFunction(inputNumber, "inputNumber")
	env.AddProcedureFromFunction(inputString, "inputString")
	env.AddProcedureFromFunction(inputString, "readLine")
	env.AddProcedureFromFunction(toSymbol, "toSymbol")

	//Lists
	env.AddProcedureFromFunction(toList, "list")
	env.AddProcedureFromFunction(car, "car")
	env.AddProcedureFromFunction(cdr, "cdr")
	env.AddProcedureFromFunction(cons, "cons")
	env.AddProcedureFromFunction(lisplen, "len")
	env.AddProcedureFromFunction(concatLists, "++")
	env.AddProcedureFromFunction(is_empty, "empty?")

	//Logic
	env.AddProcedureFromFunction(gt, ">")
	env.AddProcedureFromFunction(lt, "<")
	env.AddProcedureFromFunction(ge, ">=")
	env.AddProcedureFromFunction(le, "<=")
	env.AddProcedureFromFunction(eq, "=")

	//Boolean Logic
	env.AddProcedureFromFunction(and, "and")
	env.AddProcedureFromFunction(or, "or")
	env.AddProcedureFromFunction(not, "not")

	//TypeChecks
	env.AddProcedureFromFunction(is_list, "list?")
	env.AddProcedureFromFunction(is_procedure, "procedure?")
	env.AddProcedureFromFunction(is_symbol, "symbol?")
	env.AddProcedureFromFunction(is_bool, "bool?")
	env.AddProcedureFromFunction(is_number, "number?")
	env.AddProcedureFromFunction(is_string, "string?")
	env.AddProcedureFromFunction(is_equals, "equals?")
	env.AddProcedureFromFunction(what_type, "type?")

	//Strings
	env.AddProcedureFromFunction(charList, "toCharList")
	env.AddProcedureFromFunction(toString, "toString")

}

func gt(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(a, b float64) bool { return a > b }
	return cmp(run, tokens...)
}

func lt(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(a, b float64) bool { return a < b }
	return cmp(run, tokens...)
}

func ge(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(a, b float64) bool { return a >= b }
	return cmp(run, tokens...)
}

func le(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(a, b float64) bool { return a <= b }
	return cmp(run, tokens...)
}

func eq(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(a, b float64) bool { return a == b }
	return cmp(run, tokens...)
}

func cmp(run func(a, b float64) bool, tokens ...LispTypes.LispToken) LispTypes.LispToken {

	var lastNumber float64

	for i, thisToken := range tokens {
		if value, err := LispTypes.GetNumberContent(thisToken); err == nil {
			if i == 0 {
				lastNumber = value
			} else {
				if !run(lastNumber, value) {
					return LispTypes.LispBoolean{Contents: false}
				}
				lastNumber = value
			}
		} else {
			log.Printf("\n::ERROR:: %s not a Number.", thisToken.ValueToString())
			return nil
		}
	}
	return LispTypes.LispBoolean{Contents: true}
}

func modulo(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 { return float64(int(accumulator) % int(newvalue)) }
	return aritm(run, tokens...)
}

func multiply(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 { return accumulator * newvalue }
	return aritm(run, tokens...)
}
func divide(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 { return accumulator / newvalue }
	return aritm(run, tokens...)
}
func add(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 { return accumulator + newvalue }
	return aritm(run, tokens...)
}
func sub(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 { return accumulator - newvalue }
	return aritm(run, tokens...)
}

func intpart(tokens ...LispTypes.LispToken) LispTypes.LispToken {

	if value, ok := tokens[0].(LispTypes.Number); ok {
		return LispTypes.Number{Contents: float64(int(value.Contents))}

	} else {
		log.Printf("\n::ERROR::  %s Not a Number.", tokens[0].ToString())
		return nil
	}

}

func maxnumber(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 {
		if accumulator < newvalue {
			return newvalue
		} else {
			return accumulator
		}
	}
	if len(tokens) == 1 && tokens[0].GetType() == LispTypes.LIST {
		return aritm(run, LispTypes.Unpack(tokens[0])...)
	}
	return aritm(run, tokens...)
}

func minnumber(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 {
		if accumulator > newvalue {
			return newvalue
		} else {
			return accumulator
		}
	}
	if len(tokens) == 1 && tokens[0].GetType() == LispTypes.LIST {
		return aritm(run, LispTypes.Unpack(tokens[0])...)
	}
	return aritm(run, tokens...)
}

func aritm(run func(accumulator, newvalue float64) float64, tokens ...LispTypes.LispToken) LispTypes.LispToken {
	var accumulator float64 = 1
	for i, thisToken := range tokens {

		if value, err := LispTypes.GetNumberContent(thisToken); err == nil {
			if i == 0 {
				accumulator = value
			} else {
				accumulator = run(accumulator, value)
			}
		} else {
			log.Printf("\n::ERROR:: Arithmetic error: %s not a number.", thisToken.ToString())
			return nil
		}
	}
	return LispTypes.Number{Contents: accumulator}
}

func and(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue bool) bool { return accumulator && newvalue }
	return booleanlogic(run, tokens...)
}

func or(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue bool) bool { return accumulator || newvalue }
	return booleanlogic(run, tokens...)
}

func booleanlogic(run func(accumulator, newvalue bool) bool, tokens ...LispTypes.LispToken) LispTypes.LispToken {
	var accumulator bool = false
	for i, thisToken := range tokens {

		if value, ok := thisToken.(LispTypes.LispBoolean); ok {
			if i == 0 {
				accumulator = value.Contents
			} else {
				accumulator = run(accumulator, value.Contents)
			}
		} else {
			log.Printf("\n::ERROR:: Boolean error: %s not a boolean.", thisToken.ToString())
			return nil
		}
	}
	return LispTypes.LispBoolean{Contents: accumulator}
}

func not(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) == 1 {
		if value, ok := tokens[0].(LispTypes.LispBoolean); ok {
			return LispTypes.LispBoolean{Contents: !value.Contents}
		} else {
			log.Printf("\n::ERROR:: Boolean error: %s not a boolean.", tokens[0].ToString())
			return nil
		}

	} else {
		log.Printf("\n::ERROR:: Boolean error: not only takes a single argument: Usage: not BOOLEAN")
	}
	return nil
}

func begin(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return tokens[len(tokens)-1]
}

func car(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) == 1 {
		switch value := tokens[0].(type) {
		case LispTypes.List:
			return value.Contents[0]
		}
	}
	return tokens[0]
}

func cdr(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) == 1 {
		switch value := tokens[0].(type) {
		case LispTypes.List:
			return LispTypes.List{Contents: value.Contents[1:]}
		}
	}
	return LispTypes.List{Contents: tokens[1:]}
}

func printLisp(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	for _, value := range tokens {
		if value == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(value.ValueToString())
		}
	}
	return nil
}
func printfLisp(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	parts := strings.Split(tokens[0].ValueToString(), "%a")
	tokens = tokens[1:]
	lastPart := parts[len(parts)-1]
	parts = parts[:len(parts)-1]
	if len(parts) != len(tokens) {
		log.Printf("\n::ERROR:: Template Parts and Provided Elements mismatch. Length Parts %d. Length Elements %d", len(parts), (len(tokens)))
		return nil
	}
	resultingString := ""

	for i := 0; i < len(parts); i++ {
		if tokens[i] == nil {
			resultingString = resultingString + parts[i] + "nil"
		} else {
			resultingString = resultingString + parts[i] + tokens[i].ValueToString()
		}
	}
	resultingString = resultingString + lastPart

	fmt.Printf(resultingString)
	return nil

}

func inputNumber(tokens ...LispTypes.LispToken) LispTypes.LispToken {

	var f float64

	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		log.Printf("\n::ERROR:: %s Not a number input", tokens[0].ValueToString())
		return nil
	}

	return LispTypes.Number{Contents: f}
}

func inputString(tokens ...LispTypes.LispToken) LispTypes.LispToken {

	var s string

	_, err := fmt.Scanf("%s", &s)

	if err != nil {
		log.Printf("\n::ERROR:: %s Not a string input", tokens[0].ValueToString())
		return nil
	}

	return LispTypes.LispString{Contents: s}
}

func toList(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return LispTypes.List{Contents: tokens}
}

func concatLists(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	var result []LispTypes.LispToken
	for _, element := range tokens {
		switch value := element.(type) {
		case LispTypes.List:
			result = append(result, value.Contents...)
			break
		default:
			log.Printf("\n::ERROR:: %s Not a list.", element.ValueToString())
			return nil

		}
	}
	return LispTypes.List{Contents: result}
}

func cons(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 2 {
		log.Println("::ERROR:: Bad use of 'cons' function.")
	} else if value, ok := tokens[1].(LispTypes.List); ok {
		newListContent := append([]LispTypes.LispToken{tokens[0]}, value.Contents...)
		return LispTypes.List{Contents: newListContent}
	} else {
		log.Println("::ERROR:: Bad use of 'cons' function.")
	}

	return nil
}

func lisplen(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	length := len(tokens)
	if length == 1 && tokens[0].GetType() == LispTypes.LIST {
		value, _ := tokens[0].(LispTypes.List)
		length = len(value.Contents)
	}
	return LispTypes.Number{Contents: float64(length)}
}

func is_empty(tokens ...LispTypes.LispToken) LispTypes.LispToken {

	if value, ok := tokens[0].(LispTypes.List); ok {
		if len(value.Contents) == 0 {
			return LispTypes.LispBoolean{Contents: true}
		} else {
			return LispTypes.LispBoolean{Contents: false}
		}

	} else {

		log.Printf("\n::ERROR:: %s Not a list.", tokens[0].ValueToString())
		return nil
	}
}

func is_list(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.LIST, tokens...)
}
func is_number(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.NUMBER, tokens...)
}
func is_symbol(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.SYMBOL, tokens...)
}
func is_procedure(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.PROCEDURE, tokens...)
}
func is_bool(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.BOOLEAN, tokens...)
}

func is_string(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return typeCheck(LispTypes.STRING, tokens...)
}

func is_equals(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) < 2 {
		return LispTypes.LispBoolean{Contents: false}
	}

	for _, element := range tokens[1:] {
		if element != tokens[0] {
			return LispTypes.LispBoolean{Contents: false}
		}
	}
	return LispTypes.LispBoolean{Contents: true}
}

func typeCheck(typeToCheck LispTypes.InterfaceType, tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Println("::ERROR:: Bad number of arguments for type check")
		return nil
	}

	for _, token := range tokens {
		if token.GetType() != typeToCheck {
			return LispTypes.LispBoolean{Contents: false}
		}
	}
	return LispTypes.LispBoolean{Contents: true}
}

func charList(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Println("::ERROR:: Bad use of 'toCharList' function. Not enough arguments.")
		return nil
	}

	if tokens[0].GetType() != LispTypes.STRING {
		log.Printf("\n::ERROR:: Bad use of 'toCharList' function. %s not String.\n", tokens[0].ToString())
		return nil
	}
	if value, ok := tokens[0].(LispTypes.LispString); ok {
		var char_list_temp []LispTypes.LispToken
		for _, char_element := range value.GetContent() {
			char_list_temp = append(char_list_temp, LispTypes.LispString{Contents: string(char_element)})
		}
		return toList(char_list_temp...)

	} else {
		log.Println("::ERROR:: Bad use of 'toCharList' function.")
	}
	return nil
}

func what_type(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Println("::ERROR:: Bad use of 'type?' function. Usage: type? ATOM")
		return nil
	}

	//TODO: Should it return symbol or string??
	return LispTypes.Symbol{Contents: tokens[0].GetType().ToString()}
}

func toString(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Println("::ERROR:: Bad use of 'toString' function.")
		return nil
	}
	switch value := tokens[0].(type) {
	case LispTypes.List:
		result := ""
		for _, element := range value.GetContent() {
			result = result + element.ValueToString()
		}
		return LispTypes.LispString{Contents: result}
	}
	return LispTypes.LispString{Contents: tokens[0].ValueToString()}
}

func toSymbol(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Println("::ERROR:: Bad use of 'toString' function. Can only Transform Strings to Symbol.")
		return nil
	}
	if value, ok := tokens[0].(LispTypes.LispString); ok {

		return LispTypes.Symbol{Contents: value.GetContent()}

	} else {
		log.Println("::ERROR:: Bad use of 'toString' function. Can only Transform Strings to Symbol.")
	}
	return nil

}
