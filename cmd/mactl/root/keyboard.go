package root

import (
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/cmd/mactl/root/keyboard"
)

var Keyboard = &cobra.Command{
	Use:   "keyboard",
	Short: "keyboard management",
}

func init() {
	Keyboard.AddCommand(keyboard.Shortcut)
}
