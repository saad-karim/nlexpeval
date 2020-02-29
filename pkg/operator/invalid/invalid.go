package invalid

type Invalid struct{}

func (i *Invalid) Evaluate(op1 interface{}, op2 interface{}) bool {
	return false
}
