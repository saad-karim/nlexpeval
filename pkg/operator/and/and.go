package and

type And struct{}

// func (a *And) Evaluate(op1 string, op2 string) bool {
// 	o1, err := strconv.ParseBool(op1)
// 	if err != nil {
// 		fmt.Println("op1 err: ", err)
// 	}

// 	o2, err := strconv.ParseBool(op2)
// 	if err != nil {
// 		fmt.Println("op2 err: ", err)
// 	}

// 	return o1 && o2
// }

func (a *And) Evaluate(op1 interface{}, op2 interface{}) bool {
	o1 := op1.(bool)
	o2 := op2.(bool)

	return o1 && o2
}
