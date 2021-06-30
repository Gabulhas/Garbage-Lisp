package main

import (
	"GarbageLisp/Evaluator"
	"GarbageLisp/OutputHandler"
	"GarbageLisp/Parser"
	repl "GarbageLisp/Repl"
	"fmt"
	"io/ioutil"
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
			pipeline(textFromFile(os.Stdin.Name()))
			break
		case "-load":
			repl.Loop(pipeline(textFromFile(args[1])), true)
			break

		default:
			pipeline(textFromFile(args[0]))
			break
		}
	}

}

func textFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("\n::ERROR:: %s", err)
	}

	text := string(content)
	return text
}

func pipeline(program string) *Evaluator.Evaluator {
	parsed := Parser.Parse(program)
	myEval := Evaluator.NewEval()
	result := myEval.Run(parsed)
	if finalString := OutputHandler.PrettyPrint(result); finalString != "" {
		fmt.Println(OutputHandler.PrettyPrint(result))
	}
	return myEval
}
