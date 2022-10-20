package shell

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func RunScript(scriptPath, dir string, env []string) error {
	log.Infof("running bash script: %s from %s path", scriptPath, dir)
	cmd := exec.Command(getShell(), scriptPath)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = env
	return cmd.Run()
}

func RunCmd(cmd *exec.Cmd) error {
	log.Infof("running command %s", cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func getShell() string {
	bashPath := "/bin/bash"
	return bashPath
}
