package gitr

const (
	ConfigTemplate = `
copyRepoPathCdCmdToClipboard: true
scm:
  homeDir: {{.UserHomeDir}}/scm
  hosts:
  - hostname: github.com
    provider: github
    defaultBranch: main
    clone:
      homeDir: ""
      alwaysCreDir: true
      includeHostForCreDir: true
    scheme: https
  - hostname: gitlab.com
    provider: gitlab
    defaultBranch: main
    clone:
      homeDir: ""
      alwaysCreDir: true
      includeHostForCreDir: true
    scheme: https
  - hostname: bitbucket.org
    provider: bitbucket-cloud
    defaultBranch: master
    clone:
      homeDir: ""
      alwaysCreDir: true
      includeHostForCreDir: true
    scheme: https
`
)
