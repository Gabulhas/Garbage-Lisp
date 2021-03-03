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
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	repl := flag.Bool("r", false, "repl")
	input := flag.Bool("i", false, "stdin")
	fileName := flag.String("f", "", "filename")
	flag.Parse()

	if !*repl && !*input && *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *repl {
		fmt.Println("Welcome to GarbageLisp REPL")
		Repl()
	} else if *input {
		Pipeline(TextFromFile(os.Stdin.Name()))
		return
	} else if *fileName != "" {
		Pipeline(TextFromFile(*fileName))
		return
	}
	flag.PrintDefaults()

	program := "(begin (define r 10) (* pi (* r r)))"
	Pipeline(program)
	os.Exit(1)
}

func TextFromFile(filename string) string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

func Repl() {
	reader := bufio.NewReader(os.Stdin)
	myEval := Evaluator.NewEval()
	for true {
		fmt.Print("GarbageLisp>")
		text, _ := reader.ReadString('\n')
		parsed := Parser.Parse(text)
		result := myEval.Run(parsed)
		if result != nil {
			fmt.Println(PrettyPrint(result))
		}
	}

}

func Pipeline(program string) {
	parsed := Parser.Parse(program)
	myEval := Evaluator.NewEval()
	result := myEval.Run(parsed)
	fmt.Println(PrettyPrint(result))
}

func PrettyPrint(token LispTypes.LispToken) string {
	switch value := token.(type) {
	case LispTypes.List:
		var b strings.Builder
		fmt.Fprintf(&b, "\n")

		for i, p := range value.Contents {
			fmt.Fprintf(&b, "[%s]", PrettyPrint(p))

			if i != len(value.Contents)-1 {
				fmt.Fprintf(&b, ",")
			}

			fmt.Fprintf(&b, "\n")
		}
		return b.String()

		break
	case LispTypes.Number:
		return fmt.Sprintf("%f", value.Contents)
		break

	}
	return ""
}
