package Parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
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
		log.Printf("\n::ERROR:: Couldn't Load File: %s", err)
		return nil
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
		//TODO: Change this to list
		L := LispTypes.List{Contents: []LispTypes.LispToken{}}
		for parser.preTokens[0] != ")" {
			L.Append(parser.readFromTokens())
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
		} else if value, err := ParseNumber(token); err == nil {
			return value
		} else if token == "true" {
			return LispTypes.LispBoolean{Contents: true}
		} else if token == "false" || token == "nil" {
			return LispTypes.LispBoolean{Contents: false}
		} else {
			return LispTypes.Symbol{Contents: token}
		}
	}

	return SymbolOrNumber(token)
}

func SymbolOrNumber(token string) LispTypes.LispToken {

	if value, err := ParseNumber(token); err == nil {
		return value
	}
	return LispTypes.Symbol{Contents: token}
}

func ParseNumber(token string) (LispTypes.LispToken, error) {
	if containsAnyDot(token) {
		if value, err := strconv.Atoi(token); err == nil {
			return LispTypes.NewInt(int32(value)), nil
		}
		return nil, errors.New("notNumber")
	} else {
		if value, err := strconv.ParseFloat(token, 32); err == nil {
			return LispTypes.NewFloat(value), nil
		}
		return nil, errors.New("notNumber")

	}
}

func containsAnyDot(token string) bool {
	for _, r := range token {
		if r == '.' {
			return true
		}
	}
	return false
}
