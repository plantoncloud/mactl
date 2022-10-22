package zshrc

import (
	"runtime"
	"strings"
)

func Get() string {
	var zb strings.Builder
	zb.WriteString("## brew env vars begin\n")
	addEnvVars(&zb)
	zb.WriteString("## brew env vars end\n")
	return zb.String()
}

func addEnvVars(zb *strings.Builder) {
	if runtime.GOARCH != "arm64" {
		return
	}
	zb.WriteString("\n")
	zb.WriteString("export PATH=\"$PATH:/opt/homebrew/bin\"")
	zb.WriteString("\n")
}
