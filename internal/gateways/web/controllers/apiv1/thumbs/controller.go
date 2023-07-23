package status

import (
	"os"

	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/render"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller is a controller implementation for status checks
type Controller struct {
	apiv1.BaseController
	buildInfo *build.Info
}

// NewController creates new status controller instance
func NewController(bi *build.Info) *Controller {
	return &Controller{
		buildInfo: bi,
	}
}

// GetStatus godoc
// @Summary Get Thumbnail for a given image
// @Description get status
// @ID get-status
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/thumbs [get]
func (ctrl *Controller) GetThumbnail(ctx *gin.Context) {
	TiffPath := ctx.Param("tiffPath")

	// if file does not exist or is not tiff, return error
	if TiffPath == "" || TiffPath[len(TiffPath)-4:] != "tiff" {
		render.NotFoundError(ctx, "File not found")
		return
	}

	_, err := os.Stat(TiffPath)
	if err != nil {
		if os.IsNotExist(err) {
			render.NotFoundError(ctx, "File not found")
		}
	}

	// if file exists, but thumbnail does not, create thumbnail
	render.JSONAPIPayload(ctx, http.StatusOK, &Response{
		Status:   http.StatusText(http.StatusOK),
		Build:    ctrl.buildInfo,
		TiffPath: TiffPath,
	})

}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/thumbnail/:tiffPath", ctrl.GetThumbnail)
}
