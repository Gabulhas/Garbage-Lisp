package LispTypes

type InterfaceType int

const (
	SYMBOL InterfaceType = iota
	NUMBER
	ATOM
	LIST
	EXP
	ENV
	PROCEDURE
	QUOTATION
	ASSIGNMENT
	BOOLEAN
)

func (interfaceType InterfaceType) ToString() string {

	return [...]string{"SYMBOL", "NUMBER", "ATOM", "LIST", "EXP", "ENV", "PROCEDURE", "QUOTATION", "ASSIGNMENT", "BOOLEAN"}[interfaceType]
}
