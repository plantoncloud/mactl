package shortcut

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
)

var Setup = &cobra.Command{
	Use:   "setup",
	Short: "setup shortcuts",
	Run:   setupHandler,
}

func setupHandler(cmd *cobra.Command, args []string) {
	if err := openapp.Setup(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success!")
}
