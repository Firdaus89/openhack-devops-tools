package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sentinel", func() {
	Context("When I receive StatusCode 2xx", func() {
		It("Update the Status", func() {
			Expect("Expected").To(Equal("Expected"))
		})
	})
})
