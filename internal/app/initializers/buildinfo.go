package initializers

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/build"
)

// InitializeBuildInfo creates new build.Info
func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
