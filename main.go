//TODO: Add string support

package main

import (
	"GarbageLisp/Parser"
	"fmt"
)


func main() {
	program := "(begin (define r 10) (* pi (* r r)))"
	parsed := Parser.Parse(program)
	fmt.Println(parsed.ToString())
}


