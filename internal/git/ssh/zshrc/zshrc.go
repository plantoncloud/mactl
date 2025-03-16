package zshrc

func Get() string {
	return getAliases()
}

func getAliases() string {
	return `
alias sshf="ci ${HOME}/.ssh"
alias sshfc="code ${HOME}/.ssh"
alias sshcon="code ${HOME}/.ssh/config"
`
}
