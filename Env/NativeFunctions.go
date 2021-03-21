package Env

import (
	"GarbageLisp/LispTypes"
	"fmt"
	"log"
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

	env.AddProcedureFromFunction(intpart, "intpart")

	//Other
	env.AddProcedureFromFunction(begin, "begin")
	env.AddProcedureFromFunction(printLisp, "print")
	env.AddProcedureFromFunction(inputNumber, "inputNumber")

	//Lists
	env.AddProcedureFromFunction(toList, "list")
	env.AddProcedureFromFunction(car, "car")
	env.AddProcedureFromFunction(cdr, "cons")
	env.AddProcedureFromFunction(cons, "cons")
	env.AddProcedureFromFunction(lisplen, "len")
	//Logic
	env.AddProcedureFromFunction(gt, ">")
	env.AddProcedureFromFunction(lt, "<")
	env.AddProcedureFromFunction(ge, ">=")
	env.AddProcedureFromFunction(le, "<=")
	env.AddProcedureFromFunction(eq, "=")
	//TypeChecks
	env.AddProcedureFromFunction(is_list, "list?")
	env.AddProcedureFromFunction(is_procedure, "procedure?")
	env.AddProcedureFromFunction(is_symbol, "symbol?")
	env.AddProcedureFromFunction(is_bool, "bool?")
	env.AddProcedureFromFunction(is_number, "number?")

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

func cmp(run func(a, b float64) bool, tokens ...LispTypes.LispToken) LispTypes.LispBoolean {

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
			log.Fatalf("\n::ERROR:: %s not Boolean.", thisToken.ValueToString())
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
		return LispTypes.ValueToNumber(float64(int(value.Contents)))
	} else {
		log.Fatalf("\n::ERROR::  %s Not a Number.", tokens[0].ToString())
	}

	return LispTypes.ValueToNumber(float64(int(0)))

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
			log.Fatalf("\n::ERROR:: Arithmetic error: %s not a number.", thisToken.ToString())
		}
	}
	return LispTypes.ValueToNumber(accumulator)
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
		fmt.Println(value.ValueToString())
	}
	return nil
}

func inputNumber(tokens ...LispTypes.LispToken) LispTypes.LispToken {

	var f float64

	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		log.Fatalf("\n::ERROR:: %s Not a number input", tokens[0].ValueToString())
	}

	return LispTypes.Number{Contents: f}
}

func toList(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return LispTypes.List{Contents: tokens}
}

func cons(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 2 {
		log.Fatal("::ERROR:: Bad use of 'cons' function.")
	}
	if value, ok := tokens[1].(LispTypes.List); ok {
		newListContent := append([]LispTypes.LispToken{tokens[0]}, value.Contents...)
		return LispTypes.List{Contents: newListContent}
	} else {
		log.Fatal("::ERROR:: Bad use of 'cons' function.")
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

func typeCheck(typeToCheck LispTypes.InterfaceType, tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 1 {
		log.Fatal("::ERROR:: Bad number of arguments for type check")
	}

	for _, token := range tokens {
		if token.GetType() != typeToCheck {
			return LispTypes.LispBoolean{Contents: false}
		}
	}
	return LispTypes.LispBoolean{Contents: true}
}
