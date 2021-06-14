package repl

import (
	"GarbageLisp/Evaluator"
	"GarbageLisp/OutputHandler"
	"GarbageLisp/Parser"
	"bufio"
	"fmt"
	"os"
)

//Main loop
func Loop() {
	fmt.Println("Welcome to GarbageLisp REPL.")
	reader := bufio.NewReader(os.Stdin)
	myEval := Evaluator.NewEval()
	for {
		fmt.Print("GL>")
		text, _ := reader.ReadString('\n')

		if text == "\n" {
			continue
		}

		//TODO: Change to string builder
		//TODO: Make this cleaner

		for lparen, rparen := countParentheses(text); lparen > rparen; {
			fmt.Print("...")
			temp, _ := reader.ReadString('\n')
			templparen, temprparen := countParentheses(temp)
			lparen = lparen + templparen
			rparen = rparen + temprparen
			text = text + temp

		}

		parsed := Parser.Parse(text)
		result := myEval.Run(parsed)
		if result != nil {
			fmt.Println(OutputHandler.PrettyPrint(result))
		}
	}

}

func countParentheses(text string) (int, int) {
	lparen := 0
	rparen := 0

	for _, current := range text {
		if current == '(' {

			lparen = lparen + 1
		} else if current == ')' {
			rparen = rparen + 1
		}

	}
	return lparen, rparen
}
