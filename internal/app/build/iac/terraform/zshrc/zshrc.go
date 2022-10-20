package zshrc

func Get() string {
	return getAliases()
}

func getAliases() string {
	return `
alias ti="terraform init"
alias tp="terraform plan"
alias ta="terraform apply --auto-approve"
`
}
