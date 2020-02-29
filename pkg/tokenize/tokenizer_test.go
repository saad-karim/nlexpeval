package tokenize_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/saad-karim/nlexpeval/pkg/tokenize"
)

var _ = Describe("tokenizer", func() {
	var tokenizer *tokenize.Tokenizer

	BeforeEach(func() {
		tokenizer = &tokenize.Tokenizer{}
	})

	Context("parses string", func() {
		It("returns back tokens", func() {
			tokens := tokenizer.Parse("dog in cat,dog,mouse")
			Expect(len(tokens)).To(Equal(3))
		})
	})

	FContext("parses string with parenthesis", func() {
		It("returns back tokens", func() {
			tokens := tokenizer.Parse("(dog in cat,dog,mouse) and (bird in pig,bird)")
			Expect(len(tokens)).To(Equal(3))
		})
	})
})
