package status

import (
	"fmt"
	"os"

	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/render"
	"github.com/bo-at-pleno/go-thumbs/internal/helpers"
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
	tiffPath := ctx.Param("tiffPath")

	// if file does not exist or is not tiff, return error
	if tiffPath == "" || tiffPath[len(tiffPath)-4:] != "tiff" {
		render.NotFoundError(ctx, "File not found")
		return
	}

	_, err := os.Stat(tiffPath)
	if err != nil {
		if os.IsNotExist(err) {
			render.NotFoundError(ctx, "File not found")
		}
	}

	img, err := helpers.ReadTiff(tiffPath)
	if err != nil {
		notFoundErr := fmt.Sprintf("Not a tiff file: %s", err)
		render.NotFoundError(ctx, notFoundErr)
	}

	thumbs := helpers.Thumbnail(img, helpers.ThumbnailOptions{
		Width:           100,
		Height:          100,
		Interpolation:   helpers.Bilinear,
		LowerPercentile: 1,
		UpperPercentile: 99,
	})

	// return thumbs as base64 encoded string
	data, err := helpers.ImageToBase64(*thumbs)
	if err != nil {
		render.InternalServerError(ctx, err.Error())
	}

	// if file exists, but thumbnail does not, create thumbnail
	render.JSONAPIPayload(ctx, http.StatusOK, &ThumbnailResponse{
		ThumbnailBase64: data,
		TiffPath:        tiffPath,
	})

}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/thumbnail/:tiffPath", ctrl.GetThumbnail)
}
