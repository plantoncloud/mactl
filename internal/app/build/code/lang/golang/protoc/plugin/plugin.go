package plugin

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"os/exec"
)

const (
	ProtocGenGoPluginVersion     = "v1.28"
	ProtocGenGoGrpcPluginVersion = "v1.1"
)

var (
	GoPlugin   = fmt.Sprintf("google.golang.org/protobuf/cmd/protoc-gen-go@%s", ProtocGenGoPluginVersion)
	GrpcPlugin = fmt.Sprintf("google.golang.org/grpc/cmd/protoc-gen-go-grpc@%s", ProtocGenGoGrpcPluginVersion)
)

func Setup() error {
	if err := shell.RunCmd(exec.Command("go", "install", GoPlugin)); err != nil {
		return errors.Wrap(err, "failed to install go protoc plugin")
	}
	if err := shell.RunCmd(exec.Command("go", "install", GrpcPlugin)); err != nil {
		return errors.Wrap(err, "failed to install go grpc protoc plugin")
	}
	return nil
}
