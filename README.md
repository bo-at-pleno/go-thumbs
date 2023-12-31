# GoBase

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/bo-at-pleno/go-thumbs)](https://goreportcard.com/report/github.com/bo-at-pleno/go-thumbs)
[![Go Reference](https://pkg.go.dev/badge/github.com/bo-at-pleno/go-thumbs.svg)](https://pkg.go.dev/github.com/bo-at-pleno/go-thumbs)
[![codecov](https://codecov.io/gh/wajox/gobase/branch/master/graph/badge.svg?token=0K79C2LH2K)](https://codecov.io/gh/wajox/gobase)
[![Build Status](https://travis-ci.org/wajox/gobase.svg?branch=master)](https://travis-ci.org/wajox/gobase)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

This is a simple skeleton for golang application. Inspired by development experience and updated according to github.com/golang-standards/project-layout.

## How to use?

1. Clone the repository (with git client `git clone github.com/bo-at-pleno/go-thumbs [project_name]` or use it as template on github for creating a new project)
2. replace `github.com/bo-at-pleno/go-thumbs` with `[your_pkg_name]` in the all files

## Structure

* /api - OpenAPI specs, documentation generated by swag
* /cmd - apps
* /db - database migrations and seeds
* /docs - documentation
* /internal - application sources for internal usage
* /pkg - application sources for external usage(SDK and libraries)
* /test - some stuff for testing purposes

## Commands
```sh
# install dev tools(wire, golangci-lint, swag, ginkgo)
make install-tools

# start test environment from docker-compose-test.yml
make start-docker-compose-test

# stop test environment from docker-compose-test.yml
make stop-docker-compose-test

# build application
make build

# run all tests
make test-all

# run go generate
make gen

# generate OpenAPI docs with swag
make swagger

# generate source code from .proto files
make proto

# generate dependencies with wire
make deps
```

## Create new project

## With [clonegopkg](https://github.com/wajox/clonegopkg)

```sh
# install clonegopkg
go install github.com/wajox/clonegopkg@latest

# create your project
clonegopkg clone git@github.com:wajox/gobase.git github.com/wajox/newproject

# push to your git repository
cd ~/go/src/github.com/wajox/newproject
git add .
git commit -m "init project from gobase template"
git remote add origin git@github.com:wajox/newproject.git
git push origin master

```


## Tools and packages
* gin-gonic
* ginkgo with gomega
* spf13/viper
* spf13/cobra
* envy
* zerolog
* golangci-lint
* wire
* swag
* migrate
* protoc
* jsonapi
* docker with docker-compose
