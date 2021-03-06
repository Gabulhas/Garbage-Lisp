package Env

import (
	"GarbageLisp/LispTypes"
	"fmt"
	"log"
)

func InitEnvNativeFunctions(env map[string]LispTypes.LispToken) {
	//Arithmetic
	env["*"] = ProcedureFromFunction(multiply)
	env["+"] = ProcedureFromFunction(add)
	env["-"] = ProcedureFromFunction(sub)
	env["/"] = ProcedureFromFunction(divide)

	env["max"] = ProcedureFromFunction(maxnumber)
	env["min"] = ProcedureFromFunction(minnumber)
	//Other
	env["begin"] = ProcedureFromFunction(begin)
	env["print"] = ProcedureFromFunction(printLisp)
	//Lists
	env["list"] = ProcedureFromFunction(toList)
	env["car"] = ProcedureFromFunction(car)
	env["cdr"] = ProcedureFromFunction(cdr)
	env["cons"] = ProcedureFromFunction(cons)
	env["len"] = ProcedureFromFunction(lisplen)
	//Logic
	env[">"] = ProcedureFromFunction(gt)
	env["<"] = ProcedureFromFunction(lt)
	env[">="] = ProcedureFromFunction(ge)
	env["<="] = ProcedureFromFunction(le)
	env["="] = ProcedureFromFunction(eq)
	//TypeChecks
	env["list?"] = ProcedureFromFunction(is_list)
	env["procedure?"] = ProcedureFromFunction(is_procedure)
	env["symbol?"] = ProcedureFromFunction(is_symbol)
	env["bool?"] = ProcedureFromFunction(is_bool)
	env["number?"] = ProcedureFromFunction(is_number)

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
			log.Fatal("NonBoolean")
		}
	}
	return LispTypes.LispBoolean{Contents: true}
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

func maxnumber(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	run := func(accumulator, newvalue float64) float64 {
		if accumulator < newvalue {
			return newvalue
		} else {
			return accumulator
		}
	}
	if len(tokens) == 1 && tokens[0].GetType() == LispTypes.LIST {
		return aritm(run, unpackList(tokens[0])...)
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
		return aritm(run, unpackList(tokens[0])...)
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
			log.Fatal("NonNumberProduct")
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
			break
		}
	}
	return tokens[0]
}

func cdr(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) == 1 {
		switch value := tokens[0].(type) {
		case LispTypes.List:
			return LispTypes.List{Contents: value.Contents[1:]}
			break
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

func toList(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	return LispTypes.List{Contents: tokens}
}

func cons(tokens ...LispTypes.LispToken) LispTypes.LispToken {
	if len(tokens) != 2 {
		log.Fatal("Bad use of 'cons' function.")
	}
	if value, ok := tokens[1].(LispTypes.List); ok {
		newListContent := append([]LispTypes.LispToken{tokens[0]}, value.Contents...)
		return LispTypes.List{Contents: newListContent}
	} else {
		log.Fatal("Bad use of 'cons' function.")
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
		log.Fatal("Bad number of arguments for type check")
	}

	for _, token := range tokens {
		if token.GetType() != typeToCheck {
			return LispTypes.LispBoolean{Contents: false}
		}
	}
	return LispTypes.LispBoolean{Contents: true}
}
