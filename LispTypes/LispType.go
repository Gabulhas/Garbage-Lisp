package LispTypes

type LispToken interface {
	GetType() InterfaceType
	ToString() string
	ValueToString() string
}

func Unpack(token LispToken) []LispToken {

	switch v := token.(type) {
	case List:
		return v.Contents
	case Exp:
		return Unpack(v.Contents)
	default:
		return []LispToken{token}
	}
}
