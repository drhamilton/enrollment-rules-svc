//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

const module = "github.com/drhamilton/enrollment-rules-svc"

// Generate regenerates Go code from proto/*.proto into gen/.
func Generate() error {
	if _, err := exec.LookPath("protoc"); err != nil {
		return fmt.Errorf("protoc not found. Install with: sudo apt-get install -y protobuf-compiler")
	}

	gopath, err := sh.Output("go", "env", "GOPATH")
	if err != nil {
		return err
	}
	gobin := filepath.Join(strings.TrimSpace(gopath), "bin")

	if err := ensurePlugin(gobin, "protoc-gen-go", "google.golang.org/protobuf/cmd/protoc-gen-go@latest"); err != nil {
		return err
	}
	if err := ensurePlugin(gobin, "protoc-gen-go-grpc", "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"); err != nil {
		return err
	}

	env := map[string]string{
		"PATH": gobin + string(os.PathListSeparator) + os.Getenv("PATH"),
	}
	return sh.RunWith(env, "protoc",
		"--go_out=.", "--go_opt=module="+module,
		"--go-grpc_out=.", "--go-grpc_opt=module="+module,
		"-I", "proto",
		"proto/enrollment.proto",
	)
}

func ensurePlugin(gobin, name, pkg string) error {
	if _, err := os.Stat(filepath.Join(gobin, name)); err == nil {
		return nil
	}
	return sh.Run("go", "install", pkg)
}

// Tidy runs go mod tidy.
func Tidy() error {
	return sh.Run("go", "mod", "tidy")
}

// Test runs all Go tests.
func Test() error {
	return sh.RunV("go", "test", "./...")
}
