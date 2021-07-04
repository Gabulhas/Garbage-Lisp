package LispTypes

type InterfaceType int

const (
	SYMBOL InterfaceType = iota
	NUMBER
	LIST
	EXP
	PROCEDURE
	BOOLEAN
	STRING
)

func (interfaceType InterfaceType) ToString() string {

	return [...]string{"SYMBOL", "NUMBER",  "LIST", "EXP",  "PROCEDURE", "BOOLEAN", "STRING"}[interfaceType]
}
