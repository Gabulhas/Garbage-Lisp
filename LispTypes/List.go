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
	fmt.Fprintf(&b, "(")
	for i, p := range list.Contents {
		if i == 0 {
			fmt.Fprintf(&b, "%s", p.ToString())
		} else {
			fmt.Fprintf(&b, " %s", p.ToString())
		}

	}
	fmt.Fprintf(&b, ")")
	return b.String()
}

func (list *List) Append(token LispToken) {
	list.Contents = append(list.Contents, token)
}

func (list List) ValueToString() string {
	var b strings.Builder
	fmt.Fprintf(&b, "(")
	for i, p := range list.Contents {
		if i == 0 {
			fmt.Fprintf(&b, "%s", p.ValueToString())
		} else {
			fmt.Fprintf(&b, " %s", p.ValueToString())
		}

	}
	fmt.Fprintf(&b, ")")
	return b.String()
}
