// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/dependencies"
	"github.com/bo-at-pleno/go-thumbs/internal/app/initializers"
)

// Injectors from wire.go:

func BuildApplication() (*Application, error) {
	info := initializers.InitializeBuildInfo()
	container := &dependencies.Container{
		BuildInfo: info,
	}
	engine := initializers.InitializeRouter(container)
	httpServerConfig := initializers.InitializeHTTPServerConfig(engine)
	server, err := initializers.InitializeHTTPServer(httpServerConfig)
	if err != nil {
		return nil, err
	}
	application := &Application{
		httpServer: server,
		Container:  container,
	}
	return application, nil
}
