package Parser

import (
	"GarbageLisp/LispTypes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lparenRegex = regexp.MustCompile(`\(`)
var rparenRegex = regexp.MustCompile(`\)`)

type Parser struct {
	preTokens []string
}

func NewParser(program string) *Parser {
	newParser := new(Parser)
	newParser.preTokens = tokenize(program)
	return newParser
}

func Parse(program string) LispTypes.LispToken {

	return NewParser(program).readFromTokens()
}

func tokenize(program string) []string {
	a := lparenRegex.ReplaceAllString(program, " ( ")
	b := rparenRegex.ReplaceAllString(a, " ) ")
	return strings.Fields(b)

}

func (parser *Parser) readFromTokens() LispTypes.LispToken {

	if len(parser.preTokens) == 0 {
		log.Fatal("Unexpected EOF")
	}
	token := parser.preTokens[0]
	parser.preTokens = parser.preTokens[1:]

	fmt.Println(token)

	if token == "(" {
		L := LispTypes.Exp{
			Contents: LispTypes.List{Contents: []LispTypes.LispToken{}},
		}

		for parser.preTokens[0] != ")" {
			L.AppendIfList(parser.readFromTokens())
		}
		parser.preTokens = parser.preTokens[1:]
		return L

	} else if token == ")" {
		log.Fatal("Unexpected )")
	} else {
		return atom(token)
	}
	return atom(token)
}

func atom(token string) LispTypes.Atom {
	if value, err := strconv.ParseFloat(token, 32); err == nil {
		return LispTypes.Atom{Contents:
		LispTypes.Number{Contents: float32(value)},
		}
	}
	return LispTypes.Atom{Contents:
	LispTypes.Symbol{Contents: token},
	}

}
