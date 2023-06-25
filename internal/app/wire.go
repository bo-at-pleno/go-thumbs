//go:build wireinject
// +build wireinject

package app

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/dependencies"
	"github.com/bo-at-pleno/go-thumbs/internal/app/initializers"
	"github.com/google/wire"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		initializers.InitializeBuildInfo,
		wire.Struct(new(dependencies.Container), "BuildInfo"),
		initializers.InitializeRouter,
		initializers.InitializeHTTPServerConfig,
		initializers.InitializeHTTPServer,
		wire.Struct(new(Application), "HTTPServer", "Container"),
	)

	return &Application{}, nil
}
