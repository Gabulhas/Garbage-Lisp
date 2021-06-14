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
		repl.Loop()
	} else {
		if args[0] == "-" {
			pipeline(textFromFile(os.Stdin.Name()))
		} else {
			pipeline(textFromFile(args[0]))
		}
		return
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

func pipeline(program string) {
	parsed := Parser.Parse(program)
	myEval := Evaluator.NewEval()
	result := myEval.Run(parsed)
	if finalString := OutputHandler.PrettyPrint(result); finalString != "" {
		fmt.Println(OutputHandler.PrettyPrint(result))
	}
}
