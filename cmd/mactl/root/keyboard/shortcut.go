package keyboard

import (
	"github.com/plantoncloud/mactl/cmd/mactl/root/keyboard/shortcut"
	"github.com/spf13/cobra"
)

var Shortcut = &cobra.Command{
	Use:   "shortcut",
	Short: "shortcut management",
}

func init() {
	Shortcut.AddCommand(
		shortcut.List,
		shortcut.Init,
		shortcut.Edit,
	)
}
