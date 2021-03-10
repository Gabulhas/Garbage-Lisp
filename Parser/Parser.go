package Parser

import (
	"GarbageLisp/LispTypes"
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
	var expressions []LispTypes.LispToken
	parser := NewParser(program)
	for len(parser.preTokens) != 0 {
		expressions = append(expressions, parser.readFromTokens())
	}
	if len(expressions) > 0 {
		expressions := append([]LispTypes.LispToken{LispTypes.Symbol{Contents: "begin"}}, expressions...)
		finalList := LispTypes.List{Contents: expressions}
		return finalList
	} else {
		return LispTypes.LispBoolean{Contents: false}
	}

}

func tokenize(program string) []string {
	trimmed := strings.TrimRight(program, "\n")
	a := lparenRegex.ReplaceAllString(trimmed, " ( ")
	b := rparenRegex.ReplaceAllString(a, " ) ")
	return strings.Fields(b)

}

func (parser *Parser) readFromTokens() LispTypes.LispToken {

	if len(parser.preTokens) == 0 {
		log.Fatal("Unexpected EOF")
	}
	token := parser.preTokens[0]
	parser.preTokens = parser.preTokens[1:]

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
		if value, err := strconv.ParseFloat(token, 32); err == nil {
			return LispTypes.Number{Contents: value}
		} else if token == "true" {
			return LispTypes.LispBoolean{Contents: true}
		} else if token == "false" || token == "nil" {
			return LispTypes.LispBoolean{Contents: false}
		} else {
			return LispTypes.Symbol{Contents: token}
		}
		return LispTypes.Symbol{Contents: token}
	}
	return atom(token)
}

func atom(token string) LispTypes.Atom {
	if value, err := strconv.ParseFloat(token, 32); err == nil {
		return LispTypes.Atom{Contents:
		LispTypes.Number{Contents: value},
		}
	}
	return LispTypes.Atom{Contents:
	LispTypes.Symbol{Contents: token},
	}
}
