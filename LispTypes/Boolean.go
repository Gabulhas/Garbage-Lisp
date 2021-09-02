package LispTypes

import (
	"fmt"
)

type LispBoolean struct {
	Contents bool
}

func (lispboolean LispBoolean) GetType() InterfaceType {
	return BOOLEAN
}

func (lispboolean LispBoolean) GetContent() bool {
	return lispboolean.Contents
}

func (lispboolean LispBoolean) ToString() string {
	return fmt.Sprintf("%s (%t)", lispboolean.GetType().ToString(), lispboolean.Contents)
}

func (lispboolean LispBoolean) ValueToString() string {
	return fmt.Sprintf("%t", lispboolean.Contents)
}
