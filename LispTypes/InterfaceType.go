package LispTypes

type InterfaceType int

const (
	SYMBOL InterfaceType = iota
	NUMBER
	ATOM
	LIST
	EXP
	ENV
	LPAREN
	RPAREN
)

func (interfaceType InterfaceType) ToString() string {
	return [...]string{"SYMBOL", "NUMBER", "ATOM", "LIST", "EXP", "ENV", "LPAREN", "RPAREN"}[interfaceType]
}
