package zshrc

import (
	"runtime"
	"strings"
)

func Get() string {
	var zb strings.Builder
	addEnvVars(&zb)
	return zb.String()
}

func addEnvVars(zb *strings.Builder) {
	//https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
	zb.WriteString("export USE_GKE_GCLOUD_AUTH_PLUGIN=True")
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
