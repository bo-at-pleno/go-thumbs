package dependencies

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
}
