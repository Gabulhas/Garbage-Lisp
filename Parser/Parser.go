package Parser

import (
	"GarbageLisp/LispTypes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//Not my REGEX, check MAL lisp Implementation
var parserRegex = regexp.MustCompile(`[\s,]*(~@|[\[\]{}()'\x60~^@]|"(?:\\.|[^\\"])*"?|;.*|[^\s\[\]{}('"\x60,;)]*)`)

type Parser struct {
	preTokens []string
}

func NewParser(program string) *Parser {
	newParser := new(Parser)
	newParser.preTokens = tokenize(program)
	return newParser
}

func ParseFromFile(filename string) LispTypes.LispToken {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("\n::ERROR:: Couldn't Load File: %s", err)
	}

	text := string(content)
	return Parse(text)
}

func Parse(program string) LispTypes.LispToken {
	var expressions []LispTypes.LispToken
	parser := NewParser(program)
	for len(parser.preTokens) != 0 {
		expressions = append(expressions, parser.readFromTokens())
	}
	if len(expressions) == 1 {
		return expressions[0]
	}
	if len(expressions) > 1 {
		expressions := append([]LispTypes.LispToken{LispTypes.Symbol{Contents: "all"}}, expressions...)
		finalList := LispTypes.List{Contents: expressions}
		return finalList
	} else {
		return LispTypes.LispBoolean{Contents: false}
	}

}

func tokenize(program string) []string {
	splittedTokens := parserRegex.FindAllStringSubmatch(strings.TrimRight(program, "\n"), -1)
	var uniqueSplittedTokens []string
	for i := range splittedTokens {
		uniqueSplittedTokens = append(uniqueSplittedTokens, strings.TrimSpace(splittedTokens[i][0]))
	}

	return uniqueSplittedTokens

}

func (parser *Parser) readFromTokens() LispTypes.LispToken {

	if len(parser.preTokens) == 0 {
		log.Fatal("::ERROR:: Unexpected EOF.")
	}
	token := parser.preTokens[0]
	parser.preTokens = parser.preTokens[1:]

	switch token {

	case ";":
		for parser.preTokens[0] != "\n" {
			parser.preTokens = parser.preTokens[1:]
		}
	case "(":
		L := LispTypes.Exp{
			Contents: LispTypes.List{Contents: []LispTypes.LispToken{}},
		}

		for parser.preTokens[0] != ")" {
			L.AppendIfList(parser.readFromTokens())
		}
		parser.preTokens = parser.preTokens[1:]
		return L
	case ")":
		log.Fatal("::ERROR:: Unexpected )")
		break
	default:
		if strings.HasSuffix(token, "\"") && strings.HasPrefix(token, "\"") {
			result, err := strconv.Unquote(token)
			if err != nil {
				fmt.Println(err)
				return LispTypes.LispString{Contents: token}
			}
			return LispTypes.LispString{Contents: result}
		} else if value, err := strconv.ParseFloat(token, 32); err == nil {
			return LispTypes.Number{Contents: value}
		} else if token == "true" {
			return LispTypes.LispBoolean{Contents: true}
		} else if token == "false" || token == "nil" {
			return LispTypes.LispBoolean{Contents: false}
		} else {
			return LispTypes.Symbol{Contents: token}
		}
	}

	return atom(token)
}

func atom(token string) LispTypes.Atom {
	if value, err := strconv.ParseFloat(token, 32); err == nil {
		return LispTypes.Atom{Contents: LispTypes.Number{Contents: value}}
	}
	return LispTypes.Atom{Contents: LispTypes.Symbol{Contents: token}}
}
