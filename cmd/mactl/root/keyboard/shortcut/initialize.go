package shortcut

import (
	keyboardshortcutopenapp "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "initialize keyboard shortcuts",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := keyboardshortcutopenapp.Initialize(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success!")
}
