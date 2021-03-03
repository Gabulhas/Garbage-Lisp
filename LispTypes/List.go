package LispTypes

import (
	"fmt"
	"strings"
)

type List struct {
	Contents []LispToken
}

func (list List) GetType() InterfaceType {
	return LIST
}

func (list List) GetContent() []LispToken {
	return list.Contents
}

func (list List) ToString() string {
	var b strings.Builder
	for i, p := range list.Contents {
		fmt.Fprintf(&b, "[%s]", p.ToString())
		if i != len(list.Contents)-1 {
			fmt.Fprintf(&b, ",")
		}
	}
	return b.String()
}

func (list List) Append(token LispToken) List {

	newList := List{Contents: append(list.Contents, token)}
	return newList
}

func (list List) ValueToString()  string {
	var b strings.Builder
	for i, p := range list.Contents {
		fmt.Fprintf(&b, "[%s]", p.ValueToString())
		if i != len(list.Contents)-1 {
			fmt.Fprintf(&b, ",")
		}
	}
	return b.String()
}