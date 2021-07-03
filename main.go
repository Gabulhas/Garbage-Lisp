package main

import (
	"github.com/Gabulhas/Garbage-Lisp/Evaluator"
	"github.com/Gabulhas/Garbage-Lisp/OutputHandler"
	"github.com/Gabulhas/Garbage-Lisp/Parser"
	repl "github.com/Gabulhas/Garbage-Lisp/Repl"
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		repl.Loop(&Evaluator.Evaluator{}, false)
	} else {

		switch args[0] {

		case "-":
			pipeline(os.Stdin.Name())
			break
		case "-load":
			repl.Loop(pipeline(args[1]), true)
			break

		default:
			pipeline(args[0])
			break
		}
	}

}

func pipeline(filename string) *Evaluator.Evaluator {
	parsed := Parser.ParseFromFile(filename)
	myEval := Evaluator.NewEval()
	result := myEval.Run(parsed)
	if finalString := OutputHandler.PrettyPrint(result); finalString != "" {
		fmt.Println(OutputHandler.PrettyPrint(result))
	}
	return myEval
}
