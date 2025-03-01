package generate

const (
	UnAliases = `
unalias gp 2>/dev/null
unalias gc 2>/dev/null
unalias gcp 2>/dev/null
`
	GenericAliases = `
alias sed=gsed
alias rsh="source ~/.zshrc"
alias o="open ."
alias c="zed ."
alias zshrc="cat ${HOME}/.zshrc"
alias h="history | grep -i "
alias ll="ls -alH "
alias pwdcp="pwd|pbcopy"
`
	CdAliases = `
alias desk="cd ${HOME}/Desktop"
alias dld="cd ${HOME}/Downloads"
alias scm="cd ~/scm"
`
)
