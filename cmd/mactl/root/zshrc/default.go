package zshrc

import (
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/cmd/mactl/root/zshrc/default"
)

var Default = &cobra.Command{
	Use:   "default",
	Short: "manage default zshrc",
}

func init() {
	Default.AddCommand(_default.Cre, _default.Show)
}
