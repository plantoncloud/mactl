package shortcut

import (
	keyboardshortcutopenapp "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "edit keyboard shortcuts to open apps",
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	if err := keyboardshortcutopenapp.Edit(); err != nil {
		log.Fatalf("%v", err)
	}
}
