package OutputHandler

import (
	"fmt"
	"github.com/Gabulhas/Garbage-Lisp/Config"
	"github.com/Gabulhas/Garbage-Lisp/LispTypes"
	"log"
)

func PrettyPrint(token LispTypes.LispToken) string {

	if token != nil {
		return token.ValueToString()
	}
	return ""
}

func Fatalf(format string, v ...interface{}) {
	if Config.Repl {
		fmt.Printf(format, v)
	} else {
		log.Fatalf(format, v)
	}
}

func Fatal(v ...interface{}) {

}
