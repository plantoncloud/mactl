package zshrc

import "runtime"

func GetEnvVars() string {
	if runtime.GOARCH == "arm64" {
		return `
mkdir -p ${HOME}/.nvm
export NVM_DIR="${HOME}/.nvm"
[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && . "/opt/homebrew/opt/nvm/nvm.sh"  # This loads nvm
`
	}
	return `
mkdir -p ${HOME}/.nvm
export NVM_DIR="${HOME}/.nvm"
[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
`
}
