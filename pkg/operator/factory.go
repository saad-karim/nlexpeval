package operator

import (
	"strings"

	"github.com/saad-karim/nlexpeval/pkg/operator/and"
	"github.com/saad-karim/nlexpeval/pkg/operator/in"
	"github.com/saad-karim/nlexpeval/pkg/operator/invalid"
)

type Operator interface {
	Evaluate(interface{}, interface{}) bool
}

func GetOperator(operator string) Operator {
	switch strings.ToLower(operator) {
	case "in":
		return &in.In{}
	case "and":
		return &and.And{}
	}

	return &invalid.Invalid{}
}
