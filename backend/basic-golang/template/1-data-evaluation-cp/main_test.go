package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)
// behavior & spesification

var _ = Describe("Account", func() {
	Describe("ExecuteToByteBuffer", func() {
		// example based testing
		It("returns slice of bytes", func() { // []bytes
			// inisialisasi , A1 . Given
			account := Account{
				Name:    "Tony",
				Number:  "1002321",
				Balance: 1000,
			}

			// action , A2 , When
			b, err := ExecuteToByteBuffer(account)

			// assertion , A3 , Then
			// human friendly
			Expect(err).ShouldNot(HaveOccurred()) // nil
			Expect(string(b)).To(Equal("Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."))
		})
	})
})
