package zshrc

import "strings"

func Get() string {
	var zb strings.Builder
	zb.WriteString("\n\n# git aliases begin\n\n")
	zb.WriteString(getAliases())
	zb.WriteString("\n\n# git aliases end\n\n")
	zb.WriteString("\n\n# git functions begin\n\n")
	zb.WriteString(getFunctions())
	zb.WriteString("\n\n# git functions end\n\n")
	return zb.String()
}

func getAliases() string {
	return `
#code
alias gitcon="code ${HOME}/.gitconfig"
alias gitignore="code ${HOME}/.gitignore_global"
#git
alias gs="git status"
alias gad="git add --all"
alias gpl="git pull"
alias gl="git log"
alias gb="git branch -a"
alias gbn="git checkout -b "
alias gcb="git checkout "
alias grh="git reset HEAD --hard"
`
}

func getFunctions() string {
	return `
function gp() {
    branch_name=$(git rev-parse --abbrev-ref HEAD)
    git push --set-upstream origin ${branch_name}
}

function gc() {
    msg=${*}
    gad
    git commit -am ${msg}
}

function gcp() {
    msg=${*}
    gc ${msg}
    gp
}

function gcpi() {
    msg=${*}
    gc "[skip ci] ${msg}"
    gp
}

function gcmpl() {
    git checkout $(git_main_branch)
    gpl
}
`
}
