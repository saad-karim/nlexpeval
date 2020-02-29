package eval_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/saad-karim/nlexpeval/eval"
	"github.com/saad-karim/nlexpeval/pkg/tokenize"
)

var _ = Describe("eval", func() {
	var e *eval.Eval

	BeforeEach(func() {
		e = &eval.Eval{
			Tokenizer: &tokenize.Tokenizer{},
		}
	})

	Context("in", func() {
		It("returns true if value contained in list", func() {
			result := e.Evaluate("dog in cat,dog,mouse")
			Expect(result).To(Equal(true))
		})

		It("returns false if value not contained in list", func() {
			result := e.Evaluate("bird in cat,dog,mouse")
			Expect(result).To(Equal(false))
		})
	})

	Context("and", func() {
		It("returns true if value contained in list", func() {
			result := e.Evaluate("(dog in cat,dog,mouse) and (bird in pig,bird)")
			Expect(result).To(Equal(true))
		})

		It("returns false if value not contained in list", func() {
			result := e.Evaluate("bird in cat,dog,mouse")
			Expect(result).To(Equal(false))
		})
	})
})
