package main

import (
	"GarbageLisp/Evaluator"
	"GarbageLisp/LispTypes"
	"GarbageLisp/Parser"
	"bufio"
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
		REPL()
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

//Main loop
func REPL() {
	fmt.Println("Welcome to GarbageLisp REPL.")
	reader := bufio.NewReader(os.Stdin)
	myEval := Evaluator.NewEval()
	for true {
		fmt.Print("GL>")
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		parsed := Parser.Parse(text)
		result := myEval.Run(parsed)
		if result != nil {
			fmt.Println(prettyPrint(result))
		}
	}

}

func pipeline(program string) {
	parsed := Parser.Parse(program)
	myEval := Evaluator.NewEval()
	result := myEval.Run(parsed)
	if finalString := prettyPrint(result); finalString != "" {
		fmt.Println(prettyPrint(result))
	}
}

func prettyPrint(token LispTypes.LispToken) string {

	if token != nil {
		return token.ValueToString()
	}
	return ""
}
