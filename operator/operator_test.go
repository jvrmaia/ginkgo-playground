package operator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jvrmaia/ginkgo-playground/operator"
)

var _ = Describe("Operator", func() {
	Context("Adder", func(){
		It("add 1 + 1 will return 2", func(){
			Ω(Adder(1,1)).Should(BeIdenticalTo(2))
		})
	})
})
