package zshrc

import (
	"fmt"
	"runtime"
	"strings"
)

func Get() string {
	var zb strings.Builder
	zb.WriteString("\n\n# golang env vars begin\n\n")
	zb.WriteString(getEnvVars())
	zb.WriteString("\n\n# golang env vars end\n\n")
	zb.WriteString("\n\n# golang aliases begin\n\n")
	zb.WriteString(getAliases())
	zb.WriteString("\n\n# golang aliases end\n\n")
	return zb.String()
}

func getEnvVars() string {
	defaultVars := `
export GOPATH="${HOME}/gopa"
export GOOS="darwin"
export PATH="${PATH}:${GOPATH}/bin"
`
	return fmt.Sprintf("%s\nexport GOARCH=%s", defaultVars, runtime.GOARCH)
}

func getAliases() string {
	return `
alias gopa="cd ${GOPATH}"
alias gomodcp="cat go.mod|grep module|awk '{print \$2}'|pbcopy"
`
}
