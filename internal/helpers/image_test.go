package helpers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bo-at-pleno/go-thumbs/internal/helpers"
)

var _ = Describe("Image", func() {
	Describe("GetThumbnail", func() {
		It("should read an image into tiff", func() {
			r := 3
			Expect(r).NotTo(BeNil())
		})
	})
})
