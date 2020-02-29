package tokenize_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTokenize(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tokenize Suite")
}
