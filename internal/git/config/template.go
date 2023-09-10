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
	//GlobalGitConfigTemplate contains recommendations from https://github.com/dandavison/delta#get-started
	GlobalGitConfigTemplate = `
[core]
    excludesfile = ~/.gitignore_global
    autocrlf = false
    whitespace = cr-at-eol
	pager = delta
[interactive]
    diffFilter = delta --color-only
[add.interactive]
    useBuiltin = false # required for git 2.37.0
[delta]
    navigate = true    # use n and N to move between diff sections
    light = false      # set to true if you're in a terminal w/ a light background color (e.g. the default macOS terminal)
[merge]
    conflictstyle = diff3
[diff]
    colorMoved = default
[pull]
    rebase = false
[inititialize]
    defaultBranch = main

[url "git@gitlab.com:"]
    insteadOf = https://gitlab.com/

[url "git@github.com:"]
    insteadOf = https://github.com/

{{ range $dir, $gitConfigPath := . }}
[includeIf "gitdir:{{$dir}}/"]
    path = {{$gitConfigPath}}
{{ end }}
`
)
