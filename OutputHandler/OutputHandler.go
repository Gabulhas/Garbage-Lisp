package OutputHandler

import "GarbageLisp/LispTypes"

func PrettyPrint(token LispTypes.LispToken) string {

	if token != nil {
		return token.ValueToString()
	}
	return ""
}

func PrintError(message string) {

}
