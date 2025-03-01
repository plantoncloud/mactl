package zshrc

func Get() string {
	return getAliases()
}

func getAliases() string {
	return `
alias sshf="ci ${HOME}/.ssh"
alias sshfc="zed ${HOME}/.ssh"
alias sshcon="zed ${HOME}/.ssh/config"
`
}
