//TODO: Add string support

package main

import (
	"GarbageLisp/Evaluator"
	"GarbageLisp/LispTypes"
	"GarbageLisp/Parser"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	repl := flag.Bool("r", false, "repl")
	input := flag.Bool("i", false, "stdin")
	fileName := flag.String("f", "", "filename")
	flag.Parse()

	if *repl || (!*input && *fileName == "") {
		fmt.Println("Welcome to GarbageLisp REPL.")
		REPL()
	} else if *input {
		pipeline(textFromFile(os.Stdin.Name()))
		return
	} else if *fileName != "" {
		pipeline(textFromFile(*fileName))
		return
	}
	flag.PrintDefaults()

	os.Exit(1)
}

func textFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("\n::ERROR::%s", err)
	}

	text := string(content)
	return text
}

//Main loop
func REPL() {
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
		return token.ToString()
	}
	return ""
}
