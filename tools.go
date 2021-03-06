// +build tools

// upgrade to this.
// https://github.com/izumin5210/grapi/blob/master/tools.go
package tools

// Manage tool dependencies via go.mod.
//
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// https://github.com/golang/go/issues/25922
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "golang.org/x/tools/cmd/goimports"
)
