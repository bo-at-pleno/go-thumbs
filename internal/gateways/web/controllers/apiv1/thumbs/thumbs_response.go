package status

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
)

// Response is a declaration for a status response
type Response struct {
	ID       string      `jsonapi:"primary,status"`
	Status   string      `jsonapi:"attr,status"`
	Build    *build.Info `jsonapi:"attr,build"`
	TiffPath string      `jsonapi:"attr,tiff_path"`
}
