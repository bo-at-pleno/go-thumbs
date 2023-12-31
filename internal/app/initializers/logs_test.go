package initializers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bo-at-pleno/go-thumbs/internal/app/initializers"
)

var _ = Describe("Logs", func() {
	Describe("InitializeLogs() ", func() {
		It("should initialize logs", func() {
			err := InitializeLogs()

			Expect(err).To(BeNil())
		})
	})
})
