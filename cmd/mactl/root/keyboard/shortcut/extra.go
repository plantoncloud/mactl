package shortcut

import (
	"github.com/plantoncloud/mactl/cmd/mactl/root/keyboard/shortcut/extra"
	"github.com/spf13/cobra"
)

var Extra = &cobra.Command{
	Use:   "extra",
	Short: "extra shortcut management",
}

func init() {
	Extra.AddCommand(extra.Init, extra.Edit, extra.Reload, extra.List, extra.Del)
}
