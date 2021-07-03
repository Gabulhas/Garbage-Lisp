package OutputHandler

import "github.com/Gabulhas/Garbage-Lisp/LispTypes"

func PrettyPrint(token LispTypes.LispToken) string {

	if token != nil {
		return token.ValueToString()
	}
	return ""
}

func PrintError(message string) {

}
