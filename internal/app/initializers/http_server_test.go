package initializers_test

import (
	. "github.com/bo-at-pleno/go-thumbs/internal/app/initializers"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/router"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HttpServer", func() {
	Describe("InitializeHTTPServerConfig()", func() {
		var (
			r *gin.Engine
		)

		BeforeEach(func() {
			r = router.NewRouter()
		})

		It("should initialize config for http server initializer", func() {
			cfg := InitializeHTTPServerConfig(r)

			Expect(cfg).NotTo(BeNil())
			Expect(cfg.Router).To(Equal(r))
			Expect(cfg.HTTPServerAddr).NotTo(BeEmpty())
		})
	})

	Describe("InitializeHTTPServer()", func() {
		var (
			r   *gin.Engine
			cfg *HTTPServerConfig
		)

		BeforeEach(func() {
			r = router.NewRouter()
			cfg = InitializeHTTPServerConfig(r)
		})

		It("should initialize HTTP server", func() {
			srv, err := InitializeHTTPServer(cfg)

			Expect(srv).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
