package keyboard

import (
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/cmd/mactl/root/keyboard/shortcut"
)

var Shortcut = &cobra.Command{
	Use:   "shortcut",
	Short: "shortcut management",
}

func init() {
	Shortcut.AddCommand(shortcut.List, shortcut.Setup, shortcut.Extra)
}
