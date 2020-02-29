package eval

import "github.com/saad-karim/nlexpeval/pkg/operator"

type Tokenizer interface {
	Parse(str string) []string
}

type Eval struct {
	Tokenizer Tokenizer
}

func (e *Eval) Evaluate(str string) bool {
	tokens := e.Tokenizer.Parse(str)

	op := operator.GetOperator(tokens[1])

	return op.Evaluate(tokens[0], tokens[2])
}
