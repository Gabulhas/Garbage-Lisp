package repl

import (
	"GarbageLisp/Evaluator"
	"GarbageLisp/OutputHandler"
	"GarbageLisp/Parser"
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

var exitFlag int
var text string

//Main loop
func Loop() {
	fmt.Println("Welcome to GarbageLisp REPL.")
	catchSigint()

	reader := bufio.NewReader(os.Stdin)
	myEval := Evaluator.NewEval()
	for {
		exitFlag = 0

		fmt.Print("GL>")
		text, _ = reader.ReadString('\n')

		if text == "\n" {
			continue
		}

		//TODO: Change to string builder
		//TODO: Make this cleaner
		for lparen, rparen := countParentheses(text); lparen > rparen; {

			fmt.Print("...")
			temp, _ := reader.ReadString('\n')
			if exitFlag == 1 {
				break
			}
			templparen, temprparen := countParentheses(temp)
			lparen = lparen + templparen
			rparen = rparen + temprparen
			text = text + temp

		}

		parsed := Parser.Parse(text)
		result := myEval.Run(parsed)
		if result != nil && exitFlag != 1 {
			fmt.Println(OutputHandler.PrettyPrint(result))
			text = ""
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

func catchSigint() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			if exitFlag == 1 {
				os.Exit(1)
			} else {
				exitFlag++
				fmt.Println("\n(Do you wish to exit? Press Ctrl+C again if you do.)")
				text = "\n"
			}

		}
	}()
}
