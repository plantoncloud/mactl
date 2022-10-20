package _default

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	DefaultZshrc = `
export POWERLEVEL9K_INSTANT_PROMPT=quiet

if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

export ZSH="${HOME}/.oh-my-zsh"

ZSH_THEME="powerlevel10k/powerlevel10k"
plugins=(
    git
    zsh-completions
    zsh-syntax-highlighting
    zsh-autosuggestions
    kubectl
)

set -o vi
source $ZSH/oh-my-zsh.sh
[[ ! -f ${HOME}/.p10k.zsh ]] || source ${HOME}/.p10k.zsh

mactl zshrc generate
source ${HOME}/.zshrc.mactl.generated
eval "$(mcfly init zsh)"
`
)

const DefaultFileName = ".zshrc"

func Cre() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	zshrcFileLoc := filepath.Join(homeDir, DefaultFileName)
	if err := os.WriteFile(zshrcFileLoc, []byte(DefaultZshrc), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", zshrcFileLoc)
	}
	return nil
}

func Show() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get home dir")
	}
	if err := shell.RunCmd(exec.Command("cat", filepath.Join(homeDir, DefaultFileName))); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", DefaultFileName)
	}
	return nil
}
