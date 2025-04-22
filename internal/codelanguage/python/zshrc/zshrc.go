package zshrc

import (
	"strings"
)

func Get() string {
	var zb strings.Builder
	zb.WriteString("\n\n# python env vars begin\n\n")
	zb.WriteString(getEnvVars())
	zb.WriteString("\n\n# python env vars end\n\n")
	return zb.String()
}

func getEnvVars() string {
	defaultVars := `
export PATH="${PATH}:${HOME}/Library/Python/3.8/bin"
`
	return defaultVars
}
