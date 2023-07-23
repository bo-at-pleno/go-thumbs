package status_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1/thumbs"
)

var _ = Describe("Controller", func() {
	var (
		thumbsCtrl *Controller
	)

	BeforeEach(func() {
		gin.SetMode(gin.ReleaseMode)

		info := build.NewInfo()
		thumbsCtrl = NewController(info)
	})

	It("controller should not be nil", func() {
		Expect(thumbsCtrl).NotTo(BeNil())
	})

	Describe("GetStatus()", func() {
		It("should return status", func() {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, _ = http.NewRequest("GET", "/api/v1/status", nil)

			thumbsCtrl.GetThumbnail(ctx)

			Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
		})
	})
})
