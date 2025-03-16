package zshrc

import "strings"

func Get() string {
	var zb strings.Builder
	zb.WriteString("\n\n# kubectl env vars begin\n\n")
	zb.WriteString(getEnvVars())
	zb.WriteString("\n\n# kubectl env vars end\n\n")
	zb.WriteString("\n\n# kubectl aliases begin\n\n")
	zb.WriteString(getAliases())
	zb.WriteString("\n\n# kubectl aliases end\n\n")
	zb.WriteString("\n\n# kubectl functions begin\n\n")
	zb.WriteString(getFunctions())
	zb.WriteString("\n\n# kubectl functions end\n\n")
	return zb.String()
}

func getEnvVars() string {
	return `
export KUBE_EDITOR='code --wait'
export PATH="${PATH}:${HOME}/.krew/bin"
`
}

func getAliases() string {
	return `
alias kf="code ${HOME}/.kube"
alias kcon="code ${HOME}/.kube/config"
alias g="kubectl get "
alias gy="kubectl get -o yaml "
alias gd="kubectl get deployments "
alias d="kubectl describe "
alias kdel="kubectl delete "
alias kns="kubectl get ns"
alias gpw="kubectl get pods -o wide -n "
alias gdw="kubectl get deployments -o wide -n "
alias kgsw="kubectl get services -o wide -n "
alias kn="kubectl get nodes -o wide "
alias knl="kubectl get nodes -o wide --show-labels"
alias kw="kubectl get nodes -l kubernetes.io/role=node -o wide "
alias kwl="kubectl get nodes -l kubernetes.io/role=node -o wide --show-labels"
alias km="kubectl get nodes -l kubernetes.io/role=master -o wide "
alias kml="kubectl get nodes -l kubernetes.io/role=master -o wide --show-labels"
alias kl="kubectl logs -f "
alias ke="kubectl exec -it "
`
}

func getFunctions() string {
	return `
function use() {
  kubectl config use-context ${1}
  export K8S_CURRENT_CONTEXT="${1}"
}

function kaf() {
  kubectl apply -f ${*}
}

function kdf() {
  kubectl delete -f ${*}
}

function pon() {
  kubectl get pods --all-namespaces -o wide --field-selector spec.nodeName=${*}
}
`
}
