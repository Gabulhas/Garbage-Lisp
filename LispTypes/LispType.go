package LispTypes

type LispToken interface {
	GetType() InterfaceType
	ToString() string
}


