package in

import "strings"

type In struct{}

func (i *In) Evaluate(op1 interface{}, op2 interface{}) bool {
	o1 := op1.(string)
	o2 := op2.(string)

	elements := strings.Split(o2, ",")

	for _, element := range elements {
		if o1 == element {
			return true
		}
	}

	return false
}
