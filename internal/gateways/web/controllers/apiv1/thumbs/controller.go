package status

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"os"
	"strconv"

	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/render"
	"github.com/bo-at-pleno/go-thumbs/internal/helpers"
	"github.com/gin-gonic/gin"
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

// GetThumbnail godoc
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
			render.NotFoundError(ctx, fmt.Sprintf("File path provided does not exist: %s", tiffPath))
			return
		}
	}

	img, err := helpers.ReadTiff(tiffPath)
	if err != nil {
		notFoundErr := fmt.Sprintf("Not a tiff file: %s", err)
		render.NotFoundError(ctx, notFoundErr)
		return
	}

	thumbs := helpers.Thumbnail(img, helpers.ThumbnailOptions{
		Width:           100,
		Height:          100,
		Interpolation:   helpers.Bilinear,
		LowerPercentile: 1,
		UpperPercentile: 99,
	})

	buf := bytes.NewBuffer(nil)
	// Encode the image to JPEG format.
	err = jpeg.Encode(buf, *thumbs, nil)
	if err != nil {
		render.NotFoundError(ctx, fmt.Sprintf("Error encoding thumbnail: %s", err))
		return
	}
	// Render the thumbnail
	render.ImageResult(ctx, 200, buf.Bytes(), "jpeg")
}

// GetThumbnailWithDimensions gets a thumbnail of the specified tiff file with the specified width and height
func (ctrl *Controller) GetThumbnailWithDimensions(ctx *gin.Context) {
	tiffPath := ctx.Param("tiffPath")
	width := ctx.Query("width")
	height := ctx.Query("height")

	// Get the thumbnail
	buf := bytes.NewBuffer(nil)
	w, _ := strconv.ParseInt(width, 10, 0)
	h, _ := strconv.ParseInt(height, 10, 0)

	img, err := helpers.TiffToThumbnail(tiffPath, int(w), int(h))
	if err != nil {
		render.NotFoundError(ctx, fmt.Sprintf("Error getting thumbnail: %s", err))
		return
	}

	// Encode the image to JPEG format.
	err = jpeg.Encode(buf, *img, nil)
	if err != nil {
		render.NotFoundError(ctx, fmt.Sprintf("Error encoding thumbnail: %s", err))
		return
	}
	// Render the thumbnail
	render.ImageResult(ctx, 200, buf.Bytes(), "jpeg")
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/thumbnail/:tiffPath", ctrl.GetThumbnail)
	r.GET("/api/v1/thumbnail/:tiffPath/:width/:height", ctrl.GetThumbnailWithDimensions)
}
