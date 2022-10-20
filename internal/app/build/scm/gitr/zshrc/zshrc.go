package zshrc

func Get() string {
	return getAliases()
}

func getAliases() string {
	return `
alias rem="gitr rem "
alias web="gitr web "
alias clone="gitr clone "
alias commits="gitr commits "
alias prs="gitr prs "
alias tags="gitr tags "
alias issues="gitr issues "
alias releases="gitr releases "
alias branches="gitr branches "
alias pipe="gitr pipe "
`
}
