package test_with_ginkgo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	testGinkgo "github.com/ruang-guru/playground/backend/basic-golang/test-with-ginkgo"
)

var _ = Describe("TestWithGinkgo", func() {
	It("test Sum", func() {
		var sum = testGinkgo.Sum(1, 2)
		Expect(sum).To(Equal(3))
	})

	It("test Sum with parameter of negative number", func() {
		var sum = testGinkgo.Sum(-3, -2)
		Expect(sum).To(Equal(-5))
	})
})
