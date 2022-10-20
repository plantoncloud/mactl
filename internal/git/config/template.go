package config

const (
	GitConfigTemplate = `[user]
    name = "{{.Username}}"
    email = "{{.Email}}"
{{ if ne .SshKeyPath "" }}
[core]
    sshCommand = "ssh -i {{.SshKeyPath}}"
{{ end }}
`
	GlobalGitConfigTemplate = `[core]
    excludesfile = ~/.gitignore_global
    autocrlf = false
    whitespace = cr-at-eol
[pull]
    rebase = false
[inititialize]
    defaultBranch = main

[url "git@gitlab.com:"]
    insteadOf = https://gitlab.com/
{{ range $dir, $gitConfigPath := . }}
[includeIf "gitdir:{{$dir}}/"]
    path = {{$gitConfigPath}}
{{ end }}
`
)
