package zshrc

import (
	"runtime"
	"strings"
)

func Get() string {
	var zb strings.Builder
	zb.WriteString("## gcloud env vars begin\n")
	addEnvVars(&zb)
	zb.WriteString("## gcloud env vars end\n")
	return zb.String()
}

func addEnvVars(zb *strings.Builder) {
	if runtime.GOARCH == "arm64" {
		zb.WriteString(`
source "/opt/homebrew/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/path.zsh.inc"
source "/opt/homebrew/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/completion.zsh.inc"
`)
		return
	}
	zb.WriteString(`
source "/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/path.zsh.inc"
source "/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/completion.zsh.inc"
`)
}
